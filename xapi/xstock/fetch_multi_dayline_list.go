package xstock

import (
	"context"

	"github.com/evercyan/brick/xtime"
)

// FetchMultiDaylineList 批量查询股票的实时日K线数据
func FetchMultiDaylineList(ctx context.Context, codes []string) ([]*DaylineDetail, error) {
	stockList, err := FetchStockList(ctx, codes)
	if err != nil {
		return nil, err
	}
	list := make([]*DaylineDetail, 0)
	for _, stock := range stockList {
		list = append(list, &DaylineDetail{
			Code: stock.Code,
			Date: xtime.Format(stock.TradeAt, xtime.DateOnly),
			OP:   stock.OpenPrice,
			CP:   stock.Price,
			Per:  stock.Percent,
			PCP:  stock.PreClosePrice,
			HP:   stock.HighPrice,
			LP:   stock.LowPrice,
			TC:   stock.Hands,
			TA:   stock.Amount,
			Ex:   stock.Exchange,
		})
	}
	return list, nil
}
