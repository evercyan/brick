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
	lhbListURL    = `https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_DAILYBILLBOARD_DETAILSNEW&columns=ALL&filter=%s&pageNumber=1&pageSize=100&sortTypes=1,-1&sortColumns=SECURITY_CODE,TRADE_DATE&source=WEB&client=WEB`
	lhbListFilter = `(TRADE_DATE<='%s')(TRADE_DATE>='%s')`
)

// ----------------------------------------------------------------

// FetchLHBListResp ...
type FetchLHBListResp struct {
	Version string `json:"version"`
	Result  struct {
		Pages int `json:"pages"`
		Data  []struct {
			TRADEDATE         string      `json:"TRADE_DATE"`
			DEALAMOUNTRATIO   float64     `json:"DEAL_AMOUNT_RATIO"`
			BILLBOARDDEALAMT  float64     `json:"BILLBOARD_DEAL_AMT"`
			FREEMARKETCAP     float64     `json:"FREE_MARKET_CAP"`
			EXPLAIN           string      `json:"EXPLAIN"`
			SECUCODE          string      `json:"SECUCODE"`
			SECURITYCODE      string      `json:"SECURITY_CODE"`
			CLOSEPRICE        float64     `json:"CLOSE_PRICE"`
			CHANGERATE        float64     `json:"CHANGE_RATE"`
			TURNOVERRATE      float64     `json:"TURNOVERRATE"`
			D1CLOSEADJCHRATE  interface{} `json:"D1_CLOSE_ADJCHRATE"`
			D2CLOSEADJCHRATE  interface{} `json:"D2_CLOSE_ADJCHRATE"`
			D5CLOSEADJCHRATE  interface{} `json:"D5_CLOSE_ADJCHRATE"`
			D10CLOSEADJCHRATE interface{} `json:"D10_CLOSE_ADJCHRATE"`
			SECURITYNAMEABBR  string      `json:"SECURITY_NAME_ABBR"`
			EXPLANATION       string      `json:"EXPLANATION"`
			BILLBOARDSELLAMT  float64     `json:"BILLBOARD_SELL_AMT"`
			BILLBOARDBUYAMT   float64     `json:"BILLBOARD_BUY_AMT"`
			SECURITYINNERCODE string      `json:"SECURITY_INNER_CODE"`
			BILLBOARDNETAMT   float64     `json:"BILLBOARD_NET_AMT"`
			DEALNETRATIO      float64     `json:"DEAL_NET_RATIO"`
			ACCUMAMOUNT       int64       `json:"ACCUM_AMOUNT"`
			MARKET            string      `json:"MARKET"`
			CHANGETYPE        string      `json:"CHANGE_TYPE"`
			TRADEID           int         `json:"TRADE_ID"`
			BUYSEAT           int         `json:"BUY_SEAT"`
			SELLSEAT          int         `json:"SELL_SEAT"`
			BUYRATIO          float64     `json:"BUY_RATIO"`
			SELLRATIO         float64     `json:"SELL_RATIO"`
			TRADEMARKET       string      `json:"TRADE_MARKET"`
			D30CLOSEADJCHRATE interface{} `json:"D30_CLOSE_ADJCHRATE"`
			D20CLOSEADJCHRATE interface{} `json:"D20_CLOSE_ADJCHRATE"`
			TRADEMARKETCODE   string      `json:"TRADE_MARKET_CODE"`
			BUYSEATNEW        string      `json:"BUY_SEAT_NEW"`
			SELLSEATNEW       string      `json:"SELL_SEAT_NEW"`
			SECURITYTYPECODE  string      `json:"SECURITY_TYPE_CODE"`
			SUMBUYAMT         float64     `json:"SUM_BUY_AMT"`
			SUMSELLAMT        float64     `json:"SUM_SELL_AMT"`
			NETBSAMT          float64     `json:"NET_BS_AMT"`
		} `json:"data"`
		Count int `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// FetchLHBList 查询龙虎榜列表
func FetchLHBList(ctx context.Context, date string) (*FetchLHBListResp, error) {
	filter := fmt.Sprintf(lhbListFilter, date, date)
	url := fmt.Sprintf(lhbListURL, xencoding.URLEncode(filter))
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchLHBList url: %s, response: %s", url, response.String())
	resp := &FetchLHBListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	return resp, nil
}
