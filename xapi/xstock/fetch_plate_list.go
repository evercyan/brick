package xstock

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/evercyan/brick/xencoding"
	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xlog"
	"github.com/evercyan/brick/xutil"
)

// pz 单次最多 100 个

// ...
const (
	PlateTypeIndustry = 1 // 行业板块
	PlateTypeConcept  = 2 // 概念板块
	plateListURL      = "https://push2delay.eastmoney.com/api/qt/clist/get?pn=%d&pz=100&po=1&np=3&fid=f3&fs=%s&fields=%s&fltt=2"
	plateListFields   = "f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f14,f20,f21,f22,f25,f62,f64,f65,f66,f70,f71,f72,f76,f77,f78,f82,f83,f84,f128,f140,f104,f105,f134"
)

// ...
var (
	PlateTypeMap = map[int]string{
		PlateTypeIndustry: "m:90 t:2",
		PlateTypeConcept:  "m:90 t:3",
	}
	PlateNameMap = map[int]string{
		PlateTypeIndustry: "行业板块",
		PlateTypeConcept:  "概念板块",
	}
)

// ----------------------------------------------------------------

// FetchPlateListResp ...
type FetchPlateListResp struct {
	Rc     int    `json:"rc"`
	Rt     int    `json:"rt"`
	Svr    int    `json:"svr"`
	Lt     int    `json:"lt"`
	Full   int    `json:"full"`
	Dlmkts string `json:"dlmkts"`
	Data   struct {
		Total int `json:"total"`
		Diff  []struct {
			F2   float64 `json:"f2"`   // 市价
			F3   float64 `json:"f3"`   // 涨跌幅
			F4   float64 `json:"f4"`   // 涨跌价格
			F5   float64 `json:"f5"`   // 总手数
			F6   float64 `json:"f6"`   // 成交额
			F7   float64 `json:"f7"`   // 振幅
			F8   float64 `json:"f8"`   // 换手率
			F9   float64 `json:"f9"`   // 市盈率
			F10  float64 `json:"f10"`  // 量比
			F12  string  `json:"f12"`  // 板块代码
			F14  string  `json:"f14"`  // 板块名称
			F20  float64 `json:"f20"`  // 市值
			F21  float64 `json:"f21"`  // 流通市值
			F22  float64 `json:"f22"`  // 涨速
			F25  float64 `json:"f25"`  // 今年涨幅
			F62  float64 `json:"f62"`  // 主力净流入
			F64  float64 `json:"f64"`  // 超大流入
			F65  float64 `json:"f65"`  // 超大流出
			F66  float64 `json:"f66"`  // 净超大
			F70  float64 `json:"f70"`  // 大单流入
			F71  float64 `json:"f71"`  // 大单流出
			F72  float64 `json:"f72"`  // 净大单
			F76  float64 `json:"f76"`  // 中单流入
			F77  float64 `json:"f77"`  // 中单流出
			F78  float64 `json:"f78"`  // 净中单
			F82  float64 `json:"f82"`  // 小单流入
			F83  float64 `json:"f83"`  // 小单流出
			F84  float64 `json:"f84"`  // 净小单
			F128 string  `json:"f128"` // 领涨股票名称
			F140 string  `json:"f140"` // 领涨股票代码
			F104 float64 `json:"f104"` // 板块上涨数量
			F105 float64 `json:"f105"` // 板块下跌数量
			F134 float64 `json:"f134"` // 板块股票数量
		} `json:"diff"`
	} `json:"data"`
}

// PlateDetail ...
type PlateDetail struct {
	Code          string  `json:"code"`            // 板块代码
	Name          string  `json:"name"`            // 板块名称
	Type          int     `json:"type"`            // 板块类型: 1, 行业; 2, 概念;
	Price         float64 `json:"price"`           // 当前价格
	ChangePrice   float64 `json:"change_price"`    // 涨跌价格
	Percent       float64 `json:"percent"`         // 涨跌幅
	Exchange      float64 `json:"exchange"`        // 换手率
	Amount        float64 `json:"amount"`          // 成交额
	Hands         float64 `json:"hands"`           // 总手数
	Amplitude     float64 `json:"amplitude"`       // 振幅
	VR            float64 `json:"vr"`              // 量比
	MarketValue   float64 `json:"market_value"`    // 市值
	MarketFlow    float64 `json:"market_flow"`     // 流通市值
	PE            float64 `json:"pe"`              // 市盈率
	PB            float64 `json:"pb"`              // 市净率
	LargeBuy      float64 `json:"large_buy"`       // 主力流入
	LargeSell     float64 `json:"large_sell"`      // 主力流出
	LargeChange   float64 `json:"large_change"`    // 主力净流入
	SmallBuy      float64 `json:"small_buy"`       // 散户流入
	SmallSell     float64 `json:"small_sell"`      // 散户流出
	SmallChange   float64 `json:"small_change"`    // 散户净流入
	PlateHeadName string  `json:"plate_head_name"` // 板块领涨股名称
	PlateHeadCode string  `json:"plate_head_code"` // 板块领涨股代码
	StockTotal    float64 `json:"stock_total"`     // 股票数量
	StockUp       float64 `json:"stock_up"`        // 上涨数量
	StockDown     float64 `json:"stock_down"`      // 下跌数量
	Sort          int     `json:"sort"`            // 板块排序
	SortPercent   float64 `json:"sort_percent"`    // 板块排序比例
}

// FetchPlateList 查询板块列表
func FetchPlateList(ctx context.Context, types ...int) ([]*PlateDetail, error) {
	list := make([]*PlateDetail, 0)
	for pageNum := 1; ; pageNum++ {
		// 避免请求频率过快被封 IP
		if pageNum > 0 {
			time.Sleep(SleepDuration)
		}
		batchList, err := fetchPlateList(ctx, pageNum, types...)
		if err != nil {
			return nil, err
		}
		xlog.Ctx(ctx).Infof("FetchPlateList No.%d, count: %d", pageNum, len(batchList))
		if len(batchList) == 0 {
			break
		}
		list = append(list, batchList...)
		if Debug {
			break
		}
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("未查询到板块纪录")
	}
	total := float64(len(list))
	for k, _ := range list {
		list[k].Sort = k + 1
		list[k].SortPercent = xutil.Round(float64((k+1)*100)/total, 2)
	}
	return list, nil
}

// fetchPlateList 查询板块列表
func fetchPlateList(ctx context.Context, pageNum int, types ...int) ([]*PlateDetail, error) {
	plateType := xlodash.First(types, PlateTypeIndustry)
	plateTypeText, ok := PlateTypeMap[plateType]
	if !ok {
		return nil, fmt.Errorf("无效的板块类型: %d", plateType)
	}
	list := make([]*PlateDetail, 0)
	url := fmt.Sprintf(plateListURL, pageNum, xencoding.URLEncode(plateTypeText), plateListFields)
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchPlateList url: %s, response: %s", url, response.String())
	resp := &FetchPlateListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	for _, vv := range resp.Data.Diff {
		list = append(list, &PlateDetail{
			Code:          vv.F12,
			Name:          vv.F14,
			Type:          plateType,
			Price:         vv.F2,
			ChangePrice:   vv.F4,
			Percent:       vv.F3,
			Exchange:      vv.F8,
			Amount:        vv.F6,
			Hands:         vv.F5,
			Amplitude:     vv.F7,
			VR:            vv.F10,
			MarketValue:   vv.F20,
			MarketFlow:    vv.F21,
			PE:            vv.F9,
			LargeBuy:      vv.F64 + vv.F70,
			LargeSell:     vv.F65 + vv.F71,
			LargeChange:   vv.F66 + vv.F72,
			SmallBuy:      vv.F82,
			SmallSell:     vv.F83,
			SmallChange:   vv.F84,
			PlateHeadName: vv.F128,
			PlateHeadCode: vv.F140,
			StockTotal:    vv.F134,
			StockUp:       vv.F104,
			StockDown:     vv.F105,
		})
	}
	return list, nil
}
