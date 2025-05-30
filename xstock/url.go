package xstock

import (
	"fmt"

	"github.com/evercyan/brick/xlodash"
)

// GetEMKlineURL 获取K线页面
func GetEMKlineURL(code string, bkCodes ...string) string {
	url := fmt.Sprintf(
		"https://quote.eastmoney.com/basic/h5chart-iframe.html?code=%s&market=%s&type=r",
		code, GetMarketCode(code),
	)
	if bkCode := xlodash.First(bkCodes); bkCode != "" {
		url += "&bk=" + bkCode
	}
	return url
}

// GetEMMoneyURL 获取资金页面
func GetEMMoneyURL(code string) string {
	return fmt.Sprintf("https://data.eastmoney.com/zjlx/%s.html", code)
}

// GetEMChipsURL 获取筹码页面
func GetEMChipsURL(code string) string {
	return fmt.Sprintf("https://quote.eastmoney.com/concept/%s%s.html#chart-k-cyq", GetMarketPrefix(code), code)
}

// GetEMForumURL 获取股吧页面
func GetEMForumURL(code string) string {
	return fmt.Sprintf("https://guba.eastmoney.com/list,%s.html", code)
}

// GetEMF10URL 获取F10页面
func GetEMF10URL(code string) string {
	return fmt.Sprintf("https://emweb.securities.eastmoney.com/pc_hsf10/pages/index.html?type=web&code=%s&color=b#/cpbd", code)
}

// GetEMPlateURL 获取板块页面
func GetEMPlateURL(bkCode string) string {
	return fmt.Sprintf("https://data.eastmoney.com/bkzj/%s.html", bkCode)
}
