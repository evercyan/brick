package command

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/evercyan/brick/cmd/money/config"
	"github.com/evercyan/brick/cmd/money/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xcli/xtable"
	"github.com/spf13/cobra"
)

var (
	// InterestCommand ...
	InterestCommand = &cobra.Command{
		Use:     "interest",
		Aliases: []string{"i"},
		Short:   "计算利息",
		Run: func(cmd *cobra.Command, args []string) {
			a := &config.Answers{}
			err := survey.Ask(config.Questions, a)
			if err != nil {
				xcolor.Fail(config.Error, fmt.Sprintf("输入错误: %s", err.Error()))
				return
			}
			items := [][]interface{}{
				{
					"金额", a.Capital,
				},
				{
					"利率", fmt.Sprintf("%s %.2f%%", a.RateType, a.RateValue*100),
				},
				{
					"周期", fmt.Sprintf("%d %s", a.PeriodValue, a.PeriodType),
				},
				{
					"方式", a.PayType,
				},
			}
			xcolor.Success(config.Separator)
			xtable.New(items).Style(xtable.Dashed).Border(true).Render()
			xcolor.Success(config.Separator)

			switch a.PayType {
			case config.PayTypeOnce:
				rate := a.GetRateByPeroid()
				fee := a.Capital * rate * float64(a.PeriodValue)
				xcolor.Success(config.Success, fmt.Sprintf("利息: %s", internal.FormatPrice(fee)))
			case config.PayTypeMonthEqualInterest:
				rate, peroid := a.GetMonthValue()
				monthList := internal.CalMonthEqualInterest(a.Capital, rate, peroid)
				xcolor.Success(config.Success, "等额本息")
				xtable.New(monthList).Style(xtable.Dashed).Border(true).Render()
			case config.PayTypeMonthEqualCapital:
				rate, peroid := a.GetMonthValue()
				monthList := internal.CalMonthEqualCapital(a.Capital, rate, peroid)
				xcolor.Success(config.Success, "等额本金")
				xtable.New(monthList).Style(xtable.Dashed).Border(true).Render()
			}
		},
	}
)
