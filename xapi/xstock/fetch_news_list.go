package xstock

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/evercyan/brick/xhttp"
	"github.com/evercyan/brick/xlog"
	"github.com/evercyan/brick/xtime"
	"github.com/evercyan/brick/xtype"
)

// ...
var (
	newsListURL = `https://cmsdataapi.eastmoney.com/api/infomine?code=%s&marketType=0&types=1,2&startTime=%s&endTime=%s&format=yyyy-MM-dd`
)

// ----------------------------------------------------------------

// FetchNewsListResp ...
type FetchNewsListResp struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
	Data    []struct {
		Type       int    `json:"Type"`
		Title      string `json:"Title"`
		Url        string `json:"Url"`
		Code       string `json:"Code"`
		UniqueUrl  string `json:"UniqueUrl"`
		CreateTime string `json:"CreateTime"`
		Time       string `json:"Time"`
		NpDst      string `json:"Np_dst"`
	} `json:"Data"`
}

/*
   {
       "Type": 1,
       "Title": "海立股份：股票交易异常波动",
       "Url": "http://finance.eastmoney.com/a/202508133483121465.html",
       "Code": "202508133483121465",
       "UniqueUrl": "http://finance.eastmoney.com/a/202508133483121465.html",
       "CreateTime": "2025-08-13 19:14:08",
       "Time": "2025-08-13",
       "Np_dst": "CMS"
   },
   {
       "Type": 2,
       "Title": "海立股份:海立股份关于公司B股股票交易异常波动的公告",
       "Url": "https://data.eastmoney.com/notices/detail/600619/AN202508131726844474.html",
       "Code": "AN202508131726844474",
       "UniqueUrl": "https://data.eastmoney.com/notices/detail/600619/AN202508131726844474.html",
       "CreateTime": "2025-08-14 00:00:00",
       "Time": "2025-08-14",
       "Np_dst": "NOTICE"
   },
*/

// BuyerDetail ...
type NewsDetail struct {
	Type  int       `json:"type"`  // 类型: 1, 资讯; 2, 公告;
	Cate  string    `json:"cate"`  // 类型名称
	Code  string    `json:"code"`  // 编码
	Title string    `json:"title"` // 标题
	URL   string    `json:"url"`   // 链接
	Time  time.Time `json:"time"`  // 时间
}

// FetchNewsList 查询股票资讯列表
func FetchNewsList(
	ctx context.Context, code string, begin, end time.Time,
) ([]*NewsDetail, error) {
	reqURL := fmt.Sprintf(newsListURL, code, xtime.D(begin), xtime.D(end))
	response, err := xhttp.New().Get(ctx, reqURL, xhttp.RandomHeader())
	if err != nil {
		return nil, err
	}
	xlog.Ctx(ctx).Debugf("FetchNewsList url: %s, response: %s", reqURL, response.String())
	resp := &FetchNewsListResp{}
	if err := json.Unmarshal(response.Bytes(), resp); err != nil {
		return nil, err
	}
	list := make([]*NewsDetail, 0)
	for _, v := range resp.Data {
		cate := "资讯"
		if v.Type == 2 {
			cate = "公告"
		}
		list = append(list, &NewsDetail{
			Type:  v.Type,
			Cate:  cate,
			Code:  v.Code,
			Title: v.Title,
			URL:   v.Url,
			Time:  xtype.ToTime(v.CreateTime),
		})
	}
	return list, nil
}
