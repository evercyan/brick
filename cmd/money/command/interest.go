package command

import (
	"fmt"
	"math"

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
			a := &config.Answer{}
			err := survey.Ask(config.Question, a)
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
				monthList := CalMonthEqualInterest(a.Capital, rate, peroid)
				xcolor.Success(config.Success, "等额本息")
				xtable.New(monthList).Style(xtable.Dashed).Border(true).Render()
			case config.PayTypeMonthEqualCapital:
				rate, peroid := a.GetMonthValue()
				monthList := CalMonthEqualCapital(a.Capital, rate, peroid)
				xcolor.Success(config.Success, "等额本金")
				xtable.New(monthList).Style(xtable.Dashed).Border(true).Render()
			}
		},
	}
)

// CalMonthTotal 计算等额本息(总额, 月利率, 周期(月))
func CalMonthEqualInterest(
	amount float64,
	rate float64,
	peroid int,
) []*config.MonthPayment {
	list := make([]*config.MonthPayment, 0)
	monthTotal := amount * rate * math.Pow(1+rate, float64(peroid)) / (math.Pow(1+rate, float64(peroid)) - 1)
	remainCapital := amount
	total := float64(0)
	for i := 0; i < peroid; i++ {
		monthInterest := remainCapital * rate
		monthCapital := monthTotal - monthInterest
		total = total + monthTotal
		remainCapital = remainCapital - monthCapital
		totalCapital := amount - remainCapital
		totalInterest := total - totalCapital
		list = append(list, &config.MonthPayment{
			PeroidNum:     i + 1,
			MonthTotal:    internal.FormatPrice(monthTotal),
			MonthCapital:  internal.FormatPrice(monthCapital),
			MonthInterest: internal.FormatPrice(monthInterest),
			Total:         internal.FormatPrice(total),
			TotalCapital:  internal.FormatPrice(totalCapital),
			TotalInterest: internal.FormatPrice(totalInterest),
			RemainCapital: internal.FormatPrice(remainCapital),
		})
	}
	return list
}

// capital 计算等额本金(总额, 月利率, 周期(月))
func CalMonthEqualCapital(
	capital float64,
	rate float64,
	peroid int,
) []*config.MonthPayment {
	list := make([]*config.MonthPayment, 0)
	monthCapital := capital / float64(peroid)
	remainCapital := capital
	total := float64(0)
	for i := 0; i < peroid; i++ {
		monthInterest := remainCapital * rate
		monthTotal := monthCapital + monthInterest
		remainCapital = remainCapital - monthCapital
		total = total + monthTotal
		totalCapital := capital - remainCapital
		totalInterest := total - totalCapital
		list = append(list, &config.MonthPayment{
			PeroidNum:     i + 1,
			MonthTotal:    internal.FormatPrice(monthTotal),
			MonthCapital:  internal.FormatPrice(monthCapital),
			MonthInterest: internal.FormatPrice(monthInterest),
			Total:         internal.FormatPrice(total),
			TotalCapital:  internal.FormatPrice(totalCapital),
			TotalInterest: internal.FormatPrice(totalInterest),
			RemainCapital: internal.FormatPrice(remainCapital),
		})
	}
	return list
}
