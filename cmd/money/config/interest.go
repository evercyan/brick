package config

import (
	"github.com/AlecAivazis/survey/v2"
)

var (
	Questions = []*survey.Question{
		{
			Name: "Capital",
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
					PayTypeMonthEqualInterest,
					PayTypeMonthEqualCapital,
				},
				Default: PayTypeMonthEqualInterest,
			},
		},
	}
)

// ----------------------------------------------------------------

// MonthPayment 月还款信息
type MonthPayment struct {
	PeroidNum     int    `table:"期数"`
	MonthTotal    string `table:"月还金额"`
	MonthCapital  string `table:"月还本金"`
	MonthInterest string `table:"月还利息"`
	Total         string `table:"已还金额"`
	TotalCapital  string `table:"已还本金"`
	TotalInterest string `table:"已还利息"`
	RemainCapital string `table:"剩余本金"`
}

// ----------------------------------------------------------------

// Answers ...
type Answers struct {
	Capital     float64 `survey:"Capital"`
	RateType    string  `survey:"RateType"`
	RateValue   float64 `survey:"RateValue"`
	PeriodType  string  `survey:"PeriodType"`
	PeriodValue int     `survey:"PeriodValue"`
	PayType     string  `survey:"PayType"`
}

// GetRateByPeroid 根据还款周期计算利率
func (t *Answers) GetRateByPeroid() float64 {
	if (t.PeriodType == PeriodTypeDay && t.RateType == RateTypeDay) ||
		(t.PeriodType == PeriodTypeMonth && t.RateType == RateTypeMonth) ||
		(t.PeriodType == PeriodTypeYear && t.RateType == RateTypeYear) {
		return t.RateValue
	}
	switch t.PeriodType {
	case PeriodTypeDay:
		if t.RateType == RateTypeMonth {
			return t.RateValue / 30
		}
		return t.RateValue / 365
	case PeriodTypeMonth:
		if t.RateType == RateTypeDay {
			return t.RateValue * 30
		}
		return t.RateValue / 12
	case PeriodTypeYear:
		if t.RateType == RateTypeDay {
			return t.RateValue * 365
		}
		return t.RateValue * 12
	}
	return 0
}

// GetMonthValue 获取按月还款参数
func (t *Answers) GetMonthValue() (float64, int) {
	var (
		rate   float64 = 0
		peroid         = 0
	)
	switch t.PeriodType {
	case PeriodTypeDay:
		peroid = t.PeriodValue / 30
	case PeriodTypeMonth:
		peroid = t.PeriodValue
	case PeriodTypeYear:
		peroid = t.PeriodValue * 12
	}
	switch t.RateType {
	case RateTypeDay:
		rate = t.RateValue * 30
	case RateTypeMonth:
		rate = t.RateValue
	case RateTypeYear:
		rate = t.RateValue / 12
	}
	return rate, peroid
}
