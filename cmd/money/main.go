package main

import (
	"fmt"
	"math"

	"github.com/evercyan/brick/xencoding"

	"github.com/AlecAivazis/survey/v2"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xcli/xtable"
)

const (
	Error   = "💣"
	Notice  = "🐶"
	Success = "🎉"
	Time    = "🕙"
)

const (
	RateTypeDay   = "日利率"
	RateTypeMonth = "月利率"
	RateTypeYear  = "年利率"

	PeriodTypeDay   = "天"
	PeriodTypeMonth = "月"
	PeriodTypeYear  = "年"

	PayTypeOnce        = "到期一次性还清"
	PayTypeMonthEqual  = "按月还款, 等额本息"
	PayTypeMonthReduce = "按月还款, 等额本金"
)

// Answer ...
type Answer struct {
	Amount      float64 `survey:"Amount"`
	RateType    string  `survey:"RateType"`
	RateValue   float64 `survey:"RateValue"`
	PeriodType  string  `survey:"PeriodType"`
	PeriodValue int     `survey:"PeriodValue"`
	PayType     string  `survey:"PayType"`
}

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "Amount",
		Prompt: &survey.Input{
			Message: "本金多少?",
			Default: "880000",
		},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "RateType",
		Prompt: &survey.Select{
			Message: "选择利率类型:",
			Options: []string{
				RateTypeDay,
				RateTypeMonth,
				RateTypeYear,
			},
			Default: RateTypeYear,
		},
	},
	{
		Name: "RateValue",
		Prompt: &survey.Input{
			Message: "请输入利率?",
			Default: "0.0588",
		},
	},
	{
		Name: "PeriodType",
		Prompt: &survey.Select{
			Message: "选择周期类型:",
			Options: []string{
				PeriodTypeDay,
				PeriodTypeMonth,
				PeriodTypeYear,
			},
			Default: PeriodTypeYear,
		},
	},
	{
		Name: "PeriodValue",
		Prompt: &survey.Input{
			Message: "请输入周期?",
			Default: "30",
		},
	},
	{
		Name: "PayType",
		Prompt: &survey.Select{
			Message: "选择偿还方式:",
			Options: []string{
				PayTypeOnce,
				PayTypeMonthEqual,
				PayTypeMonthReduce,
			},
			Default: PayTypeMonthEqual,
		},
	},
}

// ----------------------------------------------------------------

// 月还款信息
type Repayment struct {
	QiShu         int
	YueGongZongEr float64
	YueGongBenJin float64
	YueGongLiXi   float64
	YiHuanBenJin  float64
	YiHuanLiXi    float64
	YiHuanZongEr  float64
	ShengYuBenJin float64
}

func CalPayTypeMonthEqual(amount float64, rate float64, peroid int) []Repayment {
	res := make([]Repayment, 0)

	daiKuanBenJin := amount
	daiKuanQiShu := peroid

	yueLiLv := rate
	yueGongZongEr := daiKuanBenJin * yueLiLv * math.Pow(1+yueLiLv, float64(daiKuanQiShu)) / (math.Pow(1+yueLiLv, float64(daiKuanQiShu)) - 1)
	shengYuBenJin := daiKuanBenJin
	yiHuanZongEr := float64(0)

	for i := 0; i < daiKuanQiShu; i++ {
		qiShu := i + 1
		yueGongLiXi := shengYuBenJin * yueLiLv
		yueGongBenJin := yueGongZongEr - yueGongLiXi
		yiHuanZongEr = yiHuanZongEr + yueGongZongEr
		shengYuBenJin = shengYuBenJin - yueGongBenJin
		yiHuanBenJin := daiKuanBenJin - shengYuBenJin
		yiHuanLiXi := yiHuanZongEr - yiHuanBenJin

		repaymentMonth := Repayment{
			QiShu:         qiShu,
			YueGongZongEr: yueGongZongEr,
			YueGongBenJin: yueGongBenJin,
			YueGongLiXi:   yueGongLiXi,
			YiHuanBenJin:  yueGongBenJin,
			YiHuanLiXi:    yiHuanLiXi,
			YiHuanZongEr:  yiHuanZongEr,
			ShengYuBenJin: shengYuBenJin,
		}

		res = append(res, repaymentMonth)
	}

	return res
}

func CalPayTypeMonthReduce(amount float64, rate float64, peroid int) []Repayment {
	res := make([]Repayment, 0)

	yueGongBenJin := amount / float64(peroid)
	shengYuBenJin := amount
	yiHuanZongEr := float64(0)

	for i := 0; i < peroid; i++ {
		qiShu := i + 1
		yueGongLiXi := shengYuBenJin * rate
		yueGongZongEr := yueGongBenJin + yueGongLiXi
		yiHuanZongEr := yiHuanZongEr + yueGongZongEr
		shengYuBenJin = shengYuBenJin - yueGongBenJin
		yiHuanBenJin := amount - shengYuBenJin
		yiHuanLiXi := yiHuanZongEr - yiHuanBenJin

		repaymentMonth := Repayment{
			QiShu:         qiShu,
			YueGongZongEr: yueGongZongEr,
			YueGongBenJin: yueGongBenJin,
			YueGongLiXi:   yueGongLiXi,
			YiHuanBenJin:  yiHuanBenJin,
			YiHuanLiXi:    yiHuanLiXi,
			YiHuanZongEr:  yiHuanZongEr,
			ShengYuBenJin: shengYuBenJin,
		}

		res = append(res, repaymentMonth)
	}
	return res
}

// ----------------------------------------------------------------

func main() {
	r := &Answer{}
	err := survey.Ask(qs, r)
	if err != nil {
		xcolor.Fail(Error, fmt.Sprintf("解析输入错误: %s", err.Error()))
		return
	}

	list := [][]interface{}{
		{
			"金额", r.Amount,
		},
		{
			"利率", fmt.Sprintf("%s %.2f%%", r.RateType, r.RateValue*100),
		},
		{
			"周期", fmt.Sprintf("%d %s", r.PeriodValue, r.PeriodType),
		},
		{
			"分期", r.PayType,
		},
		{
			"", "",
		},
	}
	switch r.PayType {
	case PayTypeOnce:
		// 根据周期类型转换利率
		rate := getOnceRate(r)
		fee := float64(r.Amount) * rate * float64(r.PeriodValue)
		list = append(list, []interface{}{
			"利息", fmt.Sprintf("%.2f", fee),
		})
	case PayTypeMonthEqual:
		// 等额本息
		// 总额, 月利率, 月周期
		rate, peroid := getMonthParam(r)
		list := CalPayTypeMonthEqual(r.Amount, rate, peroid)
		fmt.Println(xencoding.JSONEncode(list))
	case PayTypeMonthReduce:
		// 等额本金
	}
	fmt.Println(xtable.New(list).Style(xtable.Dashed).Border(true).Text())

}

// getMonthParam 转换按月还款
func getMonthParam(r *Answer) (float64, int) {
	var (
		rate   float64 = 0
		peroid int     = 0
	)
	switch r.PeriodType {
	case PeriodTypeDay:
		peroid = r.PeriodValue / 30
	case PeriodTypeMonth:
		peroid = r.PeriodValue
	case PeriodTypeYear:
		peroid = r.PeriodValue * 12
	}
	switch r.RateType {
	case RateTypeDay:
		rate = r.RateValue * 30
	case RateTypeMonth:
		rate = r.RateValue
	case RateTypeYear:
		rate = r.RateValue / 12
	}
	return rate, peroid
}

// getOnceRate 转换利率
func getOnceRate(r *Answer) float64 {
	// 利率和周期单位相同时, 无需转换
	if (r.PeriodType == PeriodTypeDay && r.RateType == RateTypeDay) ||
		(r.PeriodType == PeriodTypeMonth && r.RateType == RateTypeMonth) ||
		(r.PeriodType == PeriodTypeYear && r.RateType == RateTypeYear) {
		return r.RateValue
	}
	switch r.PeriodType {
	case PeriodTypeDay:
		if r.RateType == RateTypeMonth {
			return r.RateValue / 30
		}
		return r.RateValue / 365
	case PeriodTypeMonth:
		if r.RateType == RateTypeDay {
			return r.RateValue * 30
		}
		return r.RateValue / 12
	case PeriodTypeYear:
		if r.RateType == RateTypeDay {
			return r.RateValue * 365
		}
		return r.RateValue * 12
	}
	return 0
}
