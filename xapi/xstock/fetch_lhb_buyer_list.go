package xstock

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/evercyan/brick/xencoding"
	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
)

// ...
const (
	lhbBuyerListURL    = `https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_BILLBOARD_DAILYDETAILSBUY&columns=ALL&filter=%s&pageNumber=1&pageSize=50&sortTypes=-1&sortColumns=BUY`
	lhbBuyerListFilter = `(TRADE_DATE='%s')(SECURITY_CODE="%s")`
)

// ----------------------------------------------------------------

// FetchLHBBuyerListResp ...
type FetchLHBBuyerListResp struct {
	Version string `json:"version"`
	Result  struct {
		Pages int `json:"pages"`
		Data  []struct {
			SECURITY_CODE               string  `json:"SECURITY_CODE"`
			SECUCODE                    string  `json:"SECUCODE"`
			TRADE_DATE                  string  `json:"TRADE_DATE"`
			OPERATEDEPT_CODE            string  `json:"OPERATEDEPT_CODE"`
			OPERATEDEPT_NAME            string  `json:"OPERATEDEPT_NAME"`
			EXPLANATION                 string  `json:"EXPLANATION"`
			CHANGE_RATE                 float64 `json:"CHANGE_RATE"`
			CLOSE_PRICE                 float64 `json:"CLOSE_PRICE"`
			ACCUM_AMOUNT                int     `json:"ACCUM_AMOUNT"`
			ACCUM_VOLUME                int     `json:"ACCUM_VOLUME"`
			BUY                         float64 `json:"BUY"`
			SELL                        float64 `json:"SELL"`
			NET                         float64 `json:"NET"`
			RISE_PROBABILITY_3DAY       float64 `json:"RISE_PROBABILITY_3DAY"`
			TOTAL_BUYER_SALESTIMES_3DAY int     `json:"TOTAL_BUYER_SALESTIMES_3DAY"`
			CHANGE_TYPE                 string  `json:"CHANGE_TYPE"`
			OPERATEDEPT_CODE_OLD        string  `json:"OPERATEDEPT_CODE_OLD"`
			TOTAL_BUYRIO                float64 `json:"TOTAL_BUYRIO"`
			TOTAL_SELLRIO               float64 `json:"TOTAL_SELLRIO"`
			TRADE_ID                    string  `json:"TRADE_ID"`
		} `json:"data"`
		Count int `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// LHBBuyer ...
type LHBBuyer struct {
	Code string  `json:"code"` // 机构代码
	Name string  `json:"name"` // 机构名称
	Buy  float64 `json:"buy"`  // 买入
	Sell float64 `json:"sell"` // 卖出
	Net  float64 `json:"net"`  // 净买入
}

// FetchLHBBuyerList 查询龙虎榜买入机构列表
func FetchLHBBuyerList(ctx context.Context, code, date string) (map[string][]*LHBBuyer, error) {
	filter := fmt.Sprintf(lhbBuyerListFilter, date, code)
	url := fmt.Sprintf(lhbBuyerListURL, xencoding.URLEncode(filter))
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchLHBBuyerList url: %s, response: %s", url, response.String())
	resp := &FetchLHBBuyerListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	if len(resp.Result.Data) == 0 {
		return nil, fmt.Errorf("%s 在 %s 未上榜哦", code, date)
	}
	res := make(map[string][]*LHBBuyer)
	for _, v := range resp.Result.Data {
		if _, ok := res[v.EXPLANATION]; !ok {
			res[v.EXPLANATION] = make([]*LHBBuyer, 0)
		}
		res[v.EXPLANATION] = append(res[v.EXPLANATION], &LHBBuyer{
			Code: v.OPERATEDEPT_CODE,
			Name: v.OPERATEDEPT_NAME,
			Buy:  v.BUY,
			Sell: v.SELL,
			Net:  v.NET,
		})
	}
	return res, nil
}
