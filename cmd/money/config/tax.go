package config

import (
	"github.com/AlecAivazis/survey/v2"
)

var (
	TaxQuestion = []*survey.Question{
		{
			Name: "Insurance",
			Prompt: &survey.Input{
				Message: "五种险个人比例, 默认 10.5%",
				Default: "0.105",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "Fund",
			Prompt: &survey.Input{
				Message: "公积金个人比例, 默认 7%",
				Default: "0.07",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "Salary",
			Prompt: &survey.Input{
				Message: "月收入",
				Default: "50000",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "Deduction",
			Prompt: &survey.Input{
				Message: "专项附加扣除, 包括租房房贷或赡养老人等",
				Default: "0",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}
)

// ----------------------------------------------------------------

// TaxAnswer ...
type TaxAnswer struct {
	Insurance float64 `survey:"Insurance"`
	Fund      float64 `survey:"Fund"`
	Salary    float64 `survey:"Salary"`
	Deduction float64 `survey:"Deduction"`
}

// ----------------------------------------------------------------

// TaxIncome 月收入信息
type TaxIncome struct {
	Month       int    `table:"月份"`
	Salary      uint64 `table:"月收入(税前)"`
	TotalSalary uint64 `table:"年收入(税前)"`
	Insurance   string `table:"月缴社保"`
	Fund        string `table:"月缴公积金"`
	Tax         string `table:"月个税"`
	Income      uint64 `table:"月收入"`
	TotalIncome uint64 `table:"年收入"`
}
