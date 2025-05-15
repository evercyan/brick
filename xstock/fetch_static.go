package xstock

import (
	"context"
	"encoding/json"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xjson"
	"github.com/evercyan/brick/xlog"
)

// ...
const (
	staticMarketTradeURL  = "https://dq.10jqka.com.cn/fuyao/market_analysis_api/chart/v1/get_chart_data?chart_key=turnover_minute"
	staticMarketChangeURL = "https://dq.10jqka.com.cn/fuyao/up_down_distribution/distribution/v2/realtime"
)

// ----------------------------------------------------------------

// fetchMarkeTradeDetailResp ...
type fetchMarkeTradeDetailResp struct {
	StatusCode int `json:"status_code"`
	Data       struct {
		Charts struct {
			Total  int    `json:"total"`
			Name   string `json:"name"`
			Header []struct {
				Val  float64 `json:"val"`
				Name string  `json:"name"`
				Key  string  `json:"key"`
			} `json:"header"`
			PointList [][]int64 `json:"point_list"`
			Mtime     string    `json:"mtime"`
			Lines     []struct {
				Name string `json:"name"`
				Key  string `json:"key"`
			} `json:"lines"`
			PointKeyList []string `json:"point_key_list"`
			Key          string   `json:"key"`
			XLabelList   []string `json:"x_label_list"`
		} `json:"charts"`
	} `json:"data"`
	StatusMsg string `json:"status_msg"`
}

// fetchMarkeTradeDetail 市场成交额分时
func fetchMarkeTradeDetail(ctx context.Context) (*fetchMarkeTradeDetailResp, error) {
	response, err := xhttp.New().Get(ctx, staticMarketTradeURL, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("fetchMarkeTradeDetail response: %s", response.String())
	resp := &fetchMarkeTradeDetailResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ----------------------------------------------------------------

// fetchMarkeChangeDetailResp ...
type fetchMarkeChangeDetailResp struct {
	StatusCode int `json:"status_code"`
	Data       struct {
		Suspend        int    `json:"suspend"`
		LastUpdateTime string `json:"last_update_time"`
		LimitDown      int    `json:"limit_down"`
		LimitUp        int    `json:"limit_up"`
		Flat           int    `json:"flat"`
		Up             int    `json:"up"`
		Down           int    `json:"down"`
		Table          []struct {
			Value int    `json:"value"`
			Key   string `json:"key"`
		} `json:"table"`
	} `json:"data"`
	StatusMsg string `json:"status_msg"`
}

// fetchMarkeChangeDetail 市场成交涨跌实时数据
func fetchMarkeChangeDetail(ctx context.Context) (*fetchMarkeChangeDetailResp, error) {
	response, err := xhttp.New().Get(ctx, staticMarketChangeURL, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("fetchMarkeChangeDetail response: %s", response.String())
	resp := &fetchMarkeChangeDetailResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// ----------------------------------------------------------------

// Static ...
type Static struct {
	AmountToday     float64 `json:"amount_today"`     // 今日成交额
	AmountYesterday float64 `json:"amount_yesterday"` // 昨日成交额
	AmountChange    float64 `json:"amount_change"`    // 较昨日变动
	AmountPredict   float64 `json:"amount_predict"`   // 今日预测额
	CountUp         int     `json:"count_up"`         // 上涨数量
	CountDown       int     `json:"count_down"`       // 下跌数量
	CountFlat       int     `json:"count_flat"`       // 平盘数量
	CountLimitUp    int     `json:"count_limit_up"`   // 涨停数量
	CountLimitDown  int     `json:"count_limit_down"` // 跌停数量
}

// FetchStatic 查询大盘交易额和涨跌情况
func FetchStatic(ctx context.Context) *Static {
	static := &Static{}
	// 交易额
	tradeDetail, err := fetchMarkeTradeDetail(ctx)
	if err != nil {
		xlog.Ctx(ctx).Errorf("FetchStatic fetchMarkeTradeDetail err: %v", err)
	} else {
		for _, v := range tradeDetail.Data.Charts.Header {
			if v.Key == "turnover" {
				static.AmountToday = v.Val
			} else if v.Key == "turnover_pre" {
				static.AmountYesterday = v.Val
			} else if v.Key == "turnover_change" {
				static.AmountChange = v.Val
			} else if v.Key == "predict_turnover" {
				static.AmountPredict = v.Val
			}
		}
	}
	// 涨跌
	changeDetail, err := fetchMarkeChangeDetail(ctx)
	if err != nil {
		xlog.Ctx(ctx).Errorf("FetchStatic fetchMarkeChangeDetail err: %v", err)
	} else {
		static.CountUp = changeDetail.Data.Up
		static.CountDown = changeDetail.Data.Down
		static.CountFlat = changeDetail.Data.Flat
		static.CountLimitUp = changeDetail.Data.LimitUp
		static.CountLimitDown = changeDetail.Data.LimitDown
	}
	xlog.Ctx(ctx).Infof("FetchStatic static: %s", xjson.Encode(static))
	return static
}
