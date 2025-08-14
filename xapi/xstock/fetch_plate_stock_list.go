package xstock

import (
	"context"
	"fmt"
	"time"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
)

// ...
const (
	plateStockListURL    = "https://push2delay.eastmoney.com/api/qt/clist/get?pn=%d&pz=100&po=1&np=3&fid=f3&fs=b:%s&fields=%s&fltt=2"
	plateStockListFields = "f2,f3,f4,f5,f6,f7,f8,f9,f10,f12,f14,f15,f16,f17,f18,f20,f21,f23,f26,f34,f35,f38,f39,f45,f58,f100,f101,f146,f102,f103,f265,f297,f62,f64,f65,f66,f70,f71,f72,f76,f77,f78,f82,f83,f84"
)

// ----------------------------------------------------------------

// FetchPlateStockList 查询板块股票列表
func FetchPlateStockList(ctx context.Context, plateCode string) ([]*StockDetail, error) {
	list := make([]*StockDetail, 0)
	for pageNum := 1; ; pageNum++ {
		// 避免请求频率过快被封 IP
		if pageNum > 0 {
			time.Sleep(SleepDuration)
		}
		batchList, err := fetchPlateStockList(ctx, plateCode, pageNum)
		if err != nil {
			return nil, err
		}
		xlog.Ctx(ctx).Infof("FetchPlateStockList No.%d, count: %d", pageNum, len(batchList))
		if len(batchList) == 0 {
			break
		}
		list = append(list, batchList...)
		if Debug {
			break
		}
	}
	if len(list) == 0 {
		return nil, fmt.Errorf("未查询到板块股票纪录")
	}
	return list, nil
}

// fetchPlateStockList 查询板块股票列表
func fetchPlateStockList(ctx context.Context, plateCode string, pageNum int) ([]*StockDetail, error) {
	url := fmt.Sprintf(plateStockListURL, pageNum, plateCode, plateStockListFields)
	response, err := xhttp.New().Get(ctx, url, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("fetchStockList url: %s, response: %s", url, response.String())
	return getStockDetails(ctx, response.Bytes(), false)
}
