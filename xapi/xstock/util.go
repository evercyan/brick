package xstock

import (
	"fmt"
	"strings"
	"time"

	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xtime"
)

// generateEMCode 生成东财请求code
func generateEMCode(codes ...string) string {
	newCodes := make([]string, 0)
	for _, code := range codes {
		if strings.Contains(code, ".") {
			newCodes = append(newCodes, code)
			continue
		}
		if strings.HasPrefix(code, "6") {
			newCodes = append(newCodes, "1."+code)
		} else {
			newCodes = append(newCodes, "0."+code)
		}
	}
	return strings.Join(newCodes, ",")
}

// FormatDate ...
func FormatDate(t time.Time) string {
	return xtime.Format(t, xtime.DateJoin)
}

// GetMarketCode ...
func GetMarketCode(code string) string {
	if strings.HasPrefix(code, "688") || strings.HasPrefix(code, "60") {
		return "1"
	}
	return "0"
}

// GetMarketPrefix ...
func GetMarketPrefix(code string) string {
	if strings.HasPrefix(code, "6") {
		return "sh"
	}
	if strings.HasPrefix(code, "0") || strings.HasPrefix(code, "3") {
		return "sz"
	}
	return ""
}

// GetSymbol ...
func GetSymbol(v float64) string {
	if v > 0 {
		return "↑"
	} else if v < 0 {
		return "↓"
	}
	return ""
}

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
