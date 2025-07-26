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
	// 请求链接
	OperatedeptListURL       = `https://datacenter-web.eastmoney.com/api/data/v1/get?reportName=RPT_RATEDEPT_RETURNT_RANKING&columns=ALL&filter=%s&pageNumber=%d&pageSize=%d&sortTypes=-1,1&sortColumns=TOTAL_BUYER_SALESTIMES_1DAY,OPERATEDEPT_CODE&source=WEB&client=WEB`
	OperatedeptFilterDetault = url.QueryEscape(`(STATISTICSCYCLE="04")`)
)

// ----------------------------------------------------------------

// FetchOperatedeptListResp ...
type FetchOperatedeptListResp struct {
	Version string `json:"version"`
	Result  struct {
		Pages int                  `json:"pages"`
		Data  []*OperatedeptDetail `json:"data"`
		Count int                  `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// OperatedeptDetail ...
type OperatedeptDetail struct {
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
}

// FetchOperatedeptList 查询机构交易列表
func FetchOperatedeptList(ctx context.Context) ([]*OperatedeptDetail, error) {
	list := make([]*OperatedeptDetail, 0)
	pageNum, pageSize := 1, 50
	for {
		operatedeptList, err := fetchOperatedeptList(ctx, pageNum, pageSize)
		if err != nil {
			return nil, err
		}
		// 避免请求频率过快被封 IP
		time.Sleep(time.Second * 1)
		xlog.Ctx(ctx).Infof("FetchOperatedeptList No.%d, count: %d", pageNum, len(operatedeptList))
		if len(operatedeptList) == 0 {
			break
		}
		list = append(list, operatedeptList...)
		pageNum++
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("not found")
	}
	return list, nil
}

// fetchOperatedeptList 查询机构交易列表
func fetchOperatedeptList(
	ctx context.Context, pageNum, pageSize int,
) ([]*OperatedeptDetail, error) {
	reqURL := fmt.Sprintf(OperatedeptListURL, OperatedeptFilterDetault, pageNum, pageSize)
	response, err := xhttp.New().Get(ctx, reqURL, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchOperatedeptList url: %s, response: %s", reqURL, response.String())
	resp := &FetchOperatedeptListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	return resp.Result.Data, nil
}
