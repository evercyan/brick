package xstock

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xlog"
	"github.com/evercyan/brick/xtype"
)

// klt 周期类型: 日行情, 101; 周行情: 102; 月行情: 103;
// fqt 复权类型: 0, 默认; 1, 前复权; 2, 后复权;(在量化投资研究中普遍采用后复权数据)

// ...
const (
	minlineListURL = "http://push2his.eastmoney.com/api/qt/stock/trends2/get?ndays=%d&secid=%s&fields1=f1,f2,f3,f4,f5,f6&fields2=f51,f52,f53,f54,f55,f56,f57,f58"
)

// Minline ...
type Minline struct {
	Code string           `json:"code"` // 股票代码
	Date string           `json:"date"` // 交易日期
	List []*MinlineDetail `json:"list"` // 分时列表
}

// MinlineDetail ...
type MinlineDetail struct {
	Time  string  `json:"time"`  // 时间点
	Price float64 `json:"price"` // 当前价格
	TC    float64 `json:"tc"`    // 成交量
	TA    float64 `json:"ta"`    // 成交额
	PA    float64 `json:"pa"`    // 分时平均价格
}

// FetchMinlineList 查询单只股票的5日分时
func FetchMinlineList(ctx context.Context, code string, days ...int) ([]*Minline, error) {
	day := xlodash.First(days, 1)
	url := fmt.Sprintf(minlineListURL, day, generateEMCode(code))
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("fetchMinlineList url: %s, response: %s", url, response.String())
	var resp struct {
		Data struct {
			Trends []string `json:"trends"`
		} `json:"data"`
	}
	if err := json.Unmarshal(response.Bytes(), &resp); err != nil {
		return nil, err
	}
	if len(resp.Data.Trends) == 0 {
		return nil, fmt.Errorf("not found")
	}
	// 2025-08-01 14:00,19.49,19.49,19.50,19.47,351,683997.00,19.590
	// 2025-08-01 14:00, 时间点
	// 19.49, 前一分钟价格
	// 19.49, 当前实时价格
	// 19.50,
	// 19.47,
	// 351, 成交量(手)
	// 683997.00, 成交额
	// 19.590, 分时平均价格
	list := make([]*Minline, 0)
	dateMap := make(map[string][]*MinlineDetail)
	for _, item := range resp.Data.Trends {
		items := strings.Split(item, ",")
		if len(items) < 8 {
			continue
		}
		dt := strings.Split(items[0], " ")
		d, t := dt[0], dt[1]
		if _, ok := dateMap[d]; !ok {
			dateMap[d] = make([]*MinlineDetail, 0)
		}
		dateMap[d] = append(dateMap[d], &MinlineDetail{
			Time:  t,
			Price: xtype.ToFloat64(items[2]),
			TC:    xtype.ToFloat64(items[5]),
			TA:    xtype.ToFloat64(items[6]),
			PA:    xtype.ToFloat64(items[7]),
		})
	}
	for date, details := range dateMap {
		list = append(list, &Minline{
			Code: code,
			Date: date,
			List: details,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Date < list[j].Date
	})
	return list, nil
}
