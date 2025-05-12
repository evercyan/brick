package xapi

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
	"github.com/evercyan/brick/xtype"
	"github.com/evercyan/brick/xutil"
)

// klt 周期类型: 日行情, 101; 周行情: 102; 月行情: 103;
// fqt 复权类型: 0, 默认; 1, 前复权; 2, 后复权;(在量化投资研究中普遍采用后复权数据)

// ...
const (
	// 请求链接
	TradeListURL = "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=%s&fields2=%s&klt=101&fqt=0&secid=%s&beg=%s&end=%s"
	// 查询字段
	TradeListFields1 = "f1,f2,f3,f4,f5,f6"
	TradeListFields2 = "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61,f116"
)

// ----------------------------------------------------------------

// FetchKlineListResp ...
type FetchKlineListResp struct {
	Data struct {
		Klines []string `json:"klines"`
	} `json:"data"`
}

// LineDetail ...
type LineDetail struct {
	Code string  `json:"code"` // 股票代码
	Date string  `json:"date"` // 交易日期
	OP   float64 `json:"op"`   // 开盘价
	CP   float64 `json:"cp" `  // 收盘价
	Per  float64 `json:"per"`  // 涨跌幅
	PCP  float64 `json:"pcp"`  // 前收盘价
	HP   float64 `json:"hp"`   // 最高价
	LP   float64 `json:"lp" `  // 最低价
	TC   float64 `json:"tc"`   // 成交量
	TA   float64 `json:"ta"`   // 成交额
	Ex   float64 `json:"ex"`   // 换手率
}

// FetchKlineList 查询单只股票的日行情
func FetchKlineList(ctx context.Context, code, begin, end string) ([]*LineDetail, error) {
	url := fmt.Sprintf(TradeListURL, TradeListFields1, TradeListFields2, GetCode(code), begin, end)
	response, err := xhttp.New().Get(ctx, url, Header)
	if err != nil {
		return nil, err
	}
	xlog.Debugf("FetchKlineList url: %s, response: %s", url, response.String())
	resp := &FetchKlineListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	if len(resp.Data.Klines) == 0 {
		return nil, fmt.Errorf("not found")
	}
	list := make([]*LineDetail, 0)
	for _, item := range resp.Data.Klines {
		items := strings.Split(item, ",")
		list = append(list, &LineDetail{
			Code: code,
			Date: items[0],
			OP:   xtype.ToFloat64(items[1]),
			CP:   xtype.ToFloat64(items[2]),
			Per:  xtype.ToFloat64(items[8]),
			PCP:  xutil.Round(xtype.ToFloat64(items[2])-xtype.ToFloat64(items[9]), 2),
			HP:   xtype.ToFloat64(items[3]),
			LP:   xtype.ToFloat64(items[4]),
			TC:   xtype.ToFloat64(items[5]),
			TA:   xtype.ToFloat64(items[6]),
			Ex:   xtype.ToFloat64(items[10]),
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Date < list[j].Date
	})
	return list, nil
}
