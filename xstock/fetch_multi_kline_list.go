package xstock

import (
	"context"

	"github.com/evercyan/brick/xtime"
)

// FetchMultiKlineList 批量查询股票的实时日K线数据
func FetchMultiKlineList(ctx context.Context, codes []string) ([]*LineDetail, error) {
	stockList, err := FetchStockList(ctx, codes)
	if err != nil {
		return nil, err
	}
	list := make([]*LineDetail, 0)
	for _, stock := range stockList {
		list = append(list, &LineDetail{
			Code: stock.Code,
			Date: xtime.Format(stock.TradeAt, xtime.DateJoin),
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
