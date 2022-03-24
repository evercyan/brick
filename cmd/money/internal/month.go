package internal

import (
	"math"

	"github.com/evercyan/brick/cmd/money/config"
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
			MonthTotal:    FormatPrice(monthTotal),
			MonthCapital:  FormatPrice(monthCapital),
			MonthInterest: FormatPrice(monthInterest),
			Total:         FormatPrice(total),
			TotalCapital:  FormatPrice(totalCapital),
			TotalInterest: FormatPrice(totalInterest),
			RemainCapital: FormatPrice(remainCapital),
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
			MonthTotal:    FormatPrice(monthTotal),
			MonthCapital:  FormatPrice(monthCapital),
			MonthInterest: FormatPrice(monthInterest),
			Total:         FormatPrice(total),
			TotalCapital:  FormatPrice(totalCapital),
			TotalInterest: FormatPrice(totalInterest),
			RemainCapital: FormatPrice(remainCapital),
		})
	}
	return list
}
