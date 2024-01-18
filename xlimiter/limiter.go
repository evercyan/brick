package xlimiter

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/time/rate"
)

// config change type
type ConfigChangeType int

const (
	ADDED ConfigChangeType = iota
	MODIFIED
	DELETED
)

// config change event
type ChangeEvent struct {
	Changes map[string]*ConfigChange
}

type ConfigChange struct {
	OldValue   string
	NewValue   string
	ChangeType ConfigChangeType
}

type Limit = rate.Limit

const label = "endpoint"
const global = "global"
const keyPrefix = "rate_"

// 约定配置
// rate_xxx  10,100  代表每秒10个, 最大可以突发100; 配置以逗号分隔
// rate_global 为全局限流量
// rate_endpoint 为某个方法的限流
type Limiters struct {
	limiters map[string]*Limiter
	sync.RWMutex

	*Limiter
}

func NewLimiters(defaultRate int, defaultBurst int) *Limiters {
	return &Limiters{
		limiters: make(map[string]*Limiter),
		Limiter:  NewLimiter(Limit(defaultRate), defaultBurst, global),
	}
}

func (lims *Limiters) Get(endpoint string) (*Limiter, bool) {
	lims.RLock()
	lim, ok := lims.limiters[endpoint]
	lims.RUnlock()
	if !ok {
		lims.AddLimiter(endpoint, float32(lims.limiter.Limit()), lims.limiter.Burst())
	}
	return lim, ok
}

func (lims *Limiters) AddLimiter(endpoint string, limit float32, burst int) {
	lims.Lock()
	defer lims.Unlock()
	l := NewLimiter(Limit(limit), burst, endpoint)
	lims.limiters[endpoint] = l
}

func (lims *Limiters) store(endpoint string) *Limiter {
	lims.Lock()
	defer lims.Unlock()
	l := NewLimiter(lims.limiter.Limit(), lims.limiter.Burst(), endpoint) // 如果没有单独设置配置, 则将其设置为与全局一致
	lims.limiters[endpoint] = l
	return l
}

func (lims *Limiters) del(endpoint string) {
	lims.Lock()
	defer lims.Unlock()
	delete(lims.limiters, endpoint)
}

func (lims *Limiters) Change(ch *ChangeEvent) {
	lims.change(ch)
}

func (lims *Limiters) change(event *ChangeEvent) {
	for key, value := range event.Changes {
		if !strings.HasPrefix(key, keyPrefix) {
			continue
		}
		k := strings.TrimLeft(key, keyPrefix)

		switch value.ChangeType {
		case DELETED:
			if k == global {
				lims.Limiter = NewLimiter(Limit(rate.Inf), math.MaxInt32, k)
			} else {
				l, _ := lims.Get(k)
				l.burstGauge.Set(0)
				l.limitGauge.Set(0)
				lims.del(k)
			}
		case ADDED, MODIFIED:
			vals := strings.Split(value.NewValue, ",")

			limitation, e := strconv.Atoi(vals[0])
			if e != nil {
				continue
			}
			var burst int
			if len(vals) == 2 {
				burst, e = strconv.Atoi(vals[1])
				if e != nil {
					continue
				}
			}
			if burst == 0 {
				burst = limitation * 10
			}
			if k == global {
				lims.Limiter = NewLimiter(Limit(limitation), burst, k)
			} else {
				lims.AddLimiter(k, float32(limitation), burst)
				//lims.limiters[k] = NewLimiter(Limit(limitation), burst, k)
			}
		default:
		}

	}
}

type Limiter struct {
	limiter        *rate.Limiter
	isLimitedGauge prometheus.Gauge   // 0,1 代表是否已被限流
	limitedCounter prometheus.Counter // 触发限流后的计数
	requestMetrics prometheus.Counter // 正常调用的计数

	limitGauge prometheus.Gauge // 限流阈值 per second
	burstGauge prometheus.Gauge // 顶峰阈值

	endpoint string
}

// var once sync.Once
var isExceedMetrics *prometheus.GaugeVec  // 是否触发限流
var limitedMetrics *prometheus.CounterVec // 被限流的数量
var requestMetrics *prometheus.CounterVec // 正常请求数
var limitGauge *prometheus.GaugeVec       // 限流阈值
var burstGauge *prometheus.GaugeVec       // 突发阈值

func init() {
	gv := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "exceed_limit",
		Help: "Rate exceed limit",
	}, []string{label})
	isExceedMetrics = gv
	prometheus.MustRegister(gv)

	lm := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "limited_count",
		Help: "Been Rate Limited count",
	}, []string{label})
	limitedMetrics = lm
	prometheus.MustRegister(lm)

	nlm := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "not_limited_request_count",
		Help: "Not Total request count",
	}, []string{label})
	requestMetrics = nlm
	prometheus.MustRegister(nlm)

	limitGv := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "limit_gauge",
		Help: "Rate limit",
	}, []string{label})
	limitGauge = limitGv
	prometheus.MustRegister(limitGv)

	burstGv := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "burst_gauge",
		Help: "Rate burst ",
	}, []string{label})
	burstGauge = burstGv
	prometheus.MustRegister(burstGauge)
}

func NewLimiter(r Limit, b int, endpoint string) (lim *Limiter) {
	limiter := Limiter{
		limiter:  rate.NewLimiter(r, b),
		endpoint: endpoint,
	}
	lbl := getLabelValues(endpoint)
	limitedGauge, err := isExceedMetrics.GetMetricWith(lbl)
	if err == nil {
		limiter.isLimitedGauge = limitedGauge
	}
	limMetrics, err := limitedMetrics.GetMetricWith(lbl)
	if err == nil {
		limiter.limitedCounter = limMetrics
	}
	limitGauge, err := limitGauge.GetMetricWith(lbl)
	if err == nil {
		limiter.limitGauge = limitGauge
	}
	limiter.limitGauge.Set(float64(r))
	limReqMetrics, err := requestMetrics.GetMetricWith(lbl)
	if err == nil {
		limiter.requestMetrics = limReqMetrics
	}

	limBurstGauge, err := burstGauge.GetMetricWith(lbl)
	if err == nil {
		limiter.burstGauge = limBurstGauge
	}
	limBurstGauge.Set(float64(b))
	return &limiter
}

func getLabelValues(funcName string) map[string]string {
	return map[string]string{label: funcName}
}

func (lim *Limiter) Allow() bool {
	if lim == nil {
		return true
	}
	if lim.limiter == nil {
		return true
	}
	allowed := lim.limiter.AllowN(time.Now(), 1)
	if allowed {
		lim.isLimitedGauge.Set(0)
		lim.requestMetrics.Inc()
	} else {
		lim.isLimitedGauge.Set(1)
		lim.limitedCounter.Inc()
	}
	return allowed
}

//
//// SetLimit is shorthand for SetLimitAt(time.Now(), newLimit).
//func (lim *Limiter) SetLimit(newLimit Limit) {
//	lim.limiter.SetLimit(newLimit)
//}
//
//func (lim *Limiter) SetLimitAt(now time.Time, newLimit Limit) {
//	lim.limiter.SetLimitAt(now, newLimit)
//}
