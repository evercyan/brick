package xstock

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
)

// STATISTICSCYCLE, 统计周期: 01, 近一月; 02, 近三月; 03, 近六月; 04, 近一年;

// ...
var (
	buyerListURL       = `https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_RATEDEPT_RETURNT_RANKING&columns=ALL&filter=%s&pageNumber=%d&pageSize=%d&sortTypes=-1,1&sortColumns=TOTAL_BUYER_SALESTIMES_1DAY,OPERATEDEPT_CODE&source=WEB&client=WEB`
	buyerFilterDetault = url.QueryEscape(`(STATISTICSCYCLE="04")`)
)

// ----------------------------------------------------------------

// FetchBuyerListResp ...
type FetchBuyerListResp struct {
	Version string `json:"version"`
	Result  struct {
		Pages int `json:"pages"`
		Data  []struct {
			OPERATEDEPTCODE           string  `json:"OPERATEDEPT_CODE"`             // 营业部code
			OPERATEDEPTNAME           string  `json:"OPERATEDEPT_NAME"`             // 营业部名称
			STATISTICSCYCLE           string  `json:"STATISTICSCYCLE"`              // 统计周期
			AVERAGEINCREASE1DAY       float64 `json:"AVERAGE_INCREASE_1DAY"`        // +1天平均涨幅
			RISEPROBABILITY1DAY       float64 `json:"RISE_PROBABILITY_1DAY"`        // +1天上涨概率
			TOTALBUYERSALESTIMES1DAY  int     `json:"TOTAL_BUYER_SALESTIMES_1DAY"`  // +1天买入次数
			AVERAGEINCREASE2DAY       float64 `json:"AVERAGE_INCREASE_2DAY"`        // +2天平均涨幅
			RISEPROBABILITY2DAY       float64 `json:"RISE_PROBABILITY_2DAY"`        // +2天上涨概率
			TOTALBUYERSALESTIMES2DAY  int     `json:"TOTAL_BUYER_SALESTIMES_2DAY"`  // +2天买入次数
			AVERAGEINCREASE3DAY       float64 `json:"AVERAGE_INCREASE_3DAY"`        // +3天平均涨幅
			RISEPROBABILITY3DAY       float64 `json:"RISE_PROBABILITY_3DAY"`        // +3天上涨概率
			TOTALBUYERSALESTIMES3DAY  int     `json:"TOTAL_BUYER_SALESTIMES_3DAY"`  // +3天买入次数
			AVERAGEINCREASE5DAY       float64 `json:"AVERAGE_INCREASE_5DAY"`        // +5天平均涨幅
			RISEPROBABILITY5DAY       float64 `json:"RISE_PROBABILITY_5DAY"`        // +5天上涨概率
			TOTALBUYERSALESTIMES5DAY  int     `json:"TOTAL_BUYER_SALESTIMES_5DAY"`  // +5天买入次数
			AVERAGEINCREASE10DAY      float64 `json:"AVERAGE_INCREASE_10DAY"`       // +10天平均涨幅
			RISEPROBABILITY10DAY      float64 `json:"RISE_PROBABILITY_10DAY"`       // +10天上涨概率
			TOTALBUYERSALESTIMES10DAY int     `json:"TOTAL_BUYER_SALESTIMES_10DAY"` // +10天买入次数
			OPERATEDEPTCODEOLD        string  `json:"OPERATEDEPT_CODE_OLD"`         // 营业部旧code
		} `json:"data"`
		Count int `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// BuyerDetail ...
type BuyerDetail struct {
	Code    string  `json:"code"`    // 营业部code
	Name    string  `json:"name"`    // 营业部名称
	Cycle   string  `json:"cycle"`   // 统计周期
	Per1    float64 `json:"per1"`    // +1天平均涨幅
	Rate1   float64 `json:"rate1"`   // +1天上涨概率
	Count1  int     `json:"count1"`  // +1天买入次数
	Per2    float64 `json:"per2"`    // +2天平均涨幅
	Rate2   float64 `json:"rate2"`   // +2天上涨概率
	Count2  int     `json:"count2"`  // +2天买入次数
	Per3    float64 `json:"per3"`    // +3天平均涨幅
	Rate3   float64 `json:"rate3"`   // +3天上涨概率
	Count3  int     `json:"count3"`  // +3天买入次数
	Per5    float64 `json:"per5"`    // +5天平均涨幅
	Rate5   float64 `json:"rate5"`   // +5天上涨概率
	Count5  int     `json:"count5"`  // +5天买入次数
	Per10   float64 `json:"per10"`   // +10天平均涨幅
	Rate10  float64 `json:"rate10"`  // +10天上涨概率
	Count10 int     `json:"count10"` // +10天买入次数
}

// FetchBuyerList 查询机构交易列表
func FetchBuyerList(ctx context.Context) ([]*BuyerDetail, error) {
	list := make([]*BuyerDetail, 0)
	pageNum, pageSize := 1, 50
	for {
		if pageNum > 1 {
			time.Sleep(SleepDuration)
		}
		buyerList, err := fetchBuyerList(ctx, pageNum, pageSize)
		if err != nil {
			return nil, err
		}
		xlog.Ctx(ctx).Infof("FetchBuyerList No.%d, count: %d", pageNum, len(buyerList))
		if len(buyerList) == 0 {
			break
		}
		list = append(list, buyerList...)
		pageNum++
		if Debug {
			break
		}
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return list, nil
}

// fetchBuyerList 查询机构交易列表
func fetchBuyerList(
	ctx context.Context, pageNum, pageSize int,
) ([]*BuyerDetail, error) {
	reqURL := fmt.Sprintf(buyerListURL, buyerFilterDetault, pageNum, pageSize)
	response, err := xhttp.New().Get(ctx, reqURL, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchBuyerList url: %s, response: %s", reqURL, response.String())
	resp := &FetchBuyerListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	list := make([]*BuyerDetail, 0)
	for _, v := range resp.Result.Data {
		list = append(list, &BuyerDetail{
			Code:    v.OPERATEDEPTCODE,
			Name:    v.OPERATEDEPTNAME,
			Cycle:   v.STATISTICSCYCLE,
			Per1:    v.AVERAGEINCREASE1DAY,
			Rate1:   v.RISEPROBABILITY1DAY,
			Count1:  v.TOTALBUYERSALESTIMES1DAY,
			Per2:    v.AVERAGEINCREASE2DAY,
			Rate2:   v.RISEPROBABILITY2DAY,
			Count2:  v.TOTALBUYERSALESTIMES2DAY,
			Per3:    v.AVERAGEINCREASE3DAY,
			Rate3:   v.RISEPROBABILITY3DAY,
			Count3:  v.TOTALBUYERSALESTIMES3DAY,
			Per5:    v.AVERAGEINCREASE5DAY,
			Rate5:   v.RISEPROBABILITY5DAY,
			Count5:  v.TOTALBUYERSALESTIMES5DAY,
			Per10:   v.AVERAGEINCREASE10DAY,
			Rate10:  v.RISEPROBABILITY10DAY,
			Count10: v.TOTALBUYERSALESTIMES10DAY,
		})
	}
	return list, nil
}
