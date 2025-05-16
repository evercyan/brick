package xstock

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
	"github.com/evercyan/brick/xtype"
	"github.com/evercyan/brick/xutil"
)

// ...
const (
	// 请求链接
	StockListURL = "https://push2.eastmoney.com/api/qt/ulist.np/get?fltt=2&secids=%s&fields=%s"
	// 查询字段
	StockListFields = "f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f14,f20,f21,f23,f26,f38,f39,f45,f58,f100,f102,f103,f265"
)

// ----------------------------------------------------------------

// FetchStockListResp ...
type FetchStockListResp struct {
	Data struct {
		Total int     `json:"total"`
		Diff  []*diff `json:"diff"`
	} `json:"data"`
}

// diff ...
type diff struct {
	F2   float64 `json:"f2"`   // 当前价格
	F3   float64 `json:"f3"`   // 涨跌幅
	F4   float64 `json:"f4"`   // 涨跌价格
	F5   float64 `json:"f5"`   // 总手数
	F6   float64 `json:"f6"`   // 成交额
	F7   float64 `json:"f7"`   // 振幅
	F8   float64 `json:"f8"`   // 换手率
	F9   float64 `json:"f9"`   // 市盈率
	F10  float64 `json:"f10"`  // 量比
	F12  string  `json:"f12"`  // 股票代码
	F14  string  `json:"f14"`  // 股票名称
	F20  int64   `json:"f20"`  // 市值
	F21  int64   `json:"f21"`  // 流通市值
	F23  float32 `json:"f23"`  // 市净率
	F26  int64   `json:"f26"`  // 上市时间
	F38  float64 `json:"f38"`  // 总股本
	F39  float64 `json:"f39"`  // 流通股
	F45  float64 `json:"f45"`  // 净利润
	F58  float64 `json:"f58"`  // 股东权益
	F100 string  `json:"f100"` // 板块名称
	F102 string  `json:"f102"` // 地区板块
	F103 string  `json:"f103"` // 标签
	F265 string  `json:"f265"` // 板块代码
	//F11  float64 `json:"f11"`  // 5分钟涨幅
	//F13  string  `json:"f13"`  // 市场
	//F24  float32 `json:"f24"`  // 60日涨跌幅
	//F33  float64 `json:"f33"`  // 委比
	//F36  float64 `json:"f36"`  // 人均持股数
	//F101 string  `json:"f101"` // 板块领涨股名称
	//F128 string `json:"f128"` // 板块领涨股代码
}

// StockDetail ...
type StockDetail struct {
	Code        string    `json:"code"`         // 股票代码
	Name        string    `json:"name"`         // 股票名称
	Price       float64   `json:"price"`        // 当前价格
	ChangePrice float64   `json:"change_price"` // 涨跌价格
	Percent     float64   `json:"percent"`      // 涨跌幅
	Exchange    float64   `json:"exchange"`     // 换手率
	Amount      float64   `json:"amount"`       // 成交额
	Hands       float64   `json:"hands"`        // 总手数
	Amplitude   float64   `json:"amplitude"`    // 振幅
	VR          float64   `json:"vr"`           // 量比
	MarketValue int64     `json:"market_value"` // 市值
	MarketFlow  int64     `json:"market_flow"`  // 流通市值
	PE          float64   `json:"pe"`           // 市盈率
	PB          float32   `json:"pb"`           // 市净率
	ROE         float64   `json:"roe"`          // ROE
	StockTotal  float64   `json:"stock_total"`  // 总股本
	StockFlow   float64   `json:"stock_flow"`   // 流通股
	PlateName   string    `json:"plate_name"`   // 板块名称
	PlateCode   string    `json:"plate_code"`   // 板块代码
	PlateArea   string    `json:"plate_area"`   // 地区板块
	Tag         string    `json:"tag"`          // 标签
	ListingAt   time.Time `json:"listing_at"`   // 上市时间
}

// FetchStockList ...
func FetchStockList(ctx context.Context, codes []string) ([]*StockDetail, error) {
	url := fmt.Sprintf(StockListURL, GetCode(codes...), StockListFields)
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchStockList url: %s, response: %s", url, response.String())
	resp := &FetchStockListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	if len(resp.Data.Diff) == 0 {
		return nil, fmt.Errorf("not found")
	}
	list := make([]*StockDetail, 0)
	for _, v := range resp.Data.Diff {
		list = append(list, &StockDetail{
			Price:       v.F2,
			Percent:     v.F3,
			ChangePrice: v.F4,
			Hands:       v.F5,
			Amount:      v.F6,
			Amplitude:   v.F7,
			Exchange:    v.F8,
			PE:          v.F9,
			VR:          v.F10,
			Code:        v.F12,
			Name:        v.F14,
			MarketValue: v.F20,
			MarketFlow:  v.F21,
			PB:          v.F23,
			ListingAt:   xtype.ToTime(fmt.Sprint(v.F26), "20060102"),
			StockTotal:  v.F38,
			StockFlow:   v.F39,
			PlateName:   v.F100,
			PlateArea:   v.F102,
			Tag:         v.F103,
			PlateCode:   v.F265,
			ROE:         xutil.Round(v.F45*100/v.F58, 2),
		})
	}
	return list, nil
}
