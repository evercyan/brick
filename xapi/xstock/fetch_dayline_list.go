package xstock

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
	daylineListURL     = "https://push2his.eastmoney.com/api/qt/stock/kline/get?fields1=%s&fields2=%s&klt=101&fqt=0&secid=%s&beg=%s&end=%s"
	daylineListFields1 = "f1,f2,f3,f5"
	daylineListFields2 = "f51,f52,f53,f54,f55,f56,f57,f58,f59,f60,f61"
)

// DaylineDetail ...
type DaylineDetail struct {
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
	Vr   float64 `json:"vr"`   // 量比
	Gap  float64 `json:"gap"`  // 振幅
}

// FetchDaylineList 查询单只股票的日行情
func FetchDaylineList(ctx context.Context, code, begin, end string) ([]*DaylineDetail, error) {
	url := fmt.Sprintf(
		daylineListURL,
		daylineListFields1,
		daylineListFields2,
		generateEMCode(code),
		begin,
		end,
	)
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Infof("FetchDaylineList url: %s, response: %s", url, response.String())
	var resp struct {
		Data struct {
			Klines []string `json:"klines"`
		} `json:"data"`
	}
	if err := json.Unmarshal(response.Bytes(), &resp); err != nil {
		return nil, err
	}
	if len(resp.Data.Klines) == 0 {
		return nil, fmt.Errorf("not found")
	}
	list := make([]*DaylineDetail, 0)
	// 2025-08-01,19.59,19.60,19.74,19.43,163364,319752955.16,1.58,0.05,0.01,6.65
	// 2025-08-01, 日期
	// 19.59, 开盘价
	// 19.60, 收盘价
	// 19.74, 最高价
	// 19.43, 最低价
	// 163364, 交易量
	// 319752955.16, 交易额
	// 1.58,
	// 0.05, // 涨跌幅
	// 0.01, // 涨跌价格
	// 6.65, // 换手率
	for _, item := range resp.Data.Klines {
		items := strings.Split(item, ",")
		if len(items) < 11 {
			continue
		}
		list = append(list, &DaylineDetail{
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
	for k, v := range list {
		// 振幅=(最高-最低)/(收盘价/(1+涨幅))
		list[k].Gap = xutil.Round((v.HP-v.LP)/(v.CP/(100+v.Per)), 2)
		// 无量比数据...
	}
	return list, nil
}
