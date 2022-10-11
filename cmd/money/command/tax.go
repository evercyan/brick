package command

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/evercyan/brick/cmd/money/config"
	"github.com/evercyan/brick/cmd/money/internal"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xcli/xtable"
	"github.com/evercyan/brick/xconvert"
	"github.com/spf13/cobra"
)

// ATTENTION: 本程序未考虑各城市的社保公积金交纳上下限, 存在一定误差

// 个税计算
// 全月应纳税所得额 = 税前收入 - 起征点(5000) - 专项扣除(三险一金) - 专项附加扣除(租房赡养等) - 其他扣除
// 缴税 = 全月应纳税所得额 * 税率 - 速算扣除

const (
	// TaxBase 个税起征点
	TaxBase = 5000
)

var (
	// TaxCommand ...
	TaxCommand = &cobra.Command{
		Use:     "tax",
		Aliases: []string{"i"},
		Short:   "计算个税",
		Run: func(cmd *cobra.Command, args []string) {
			xcolor.Success(config.Separator)
			xcolor.Success(config.Success, "五险一金默认配置")
			xcolor.Success(config.Success, "社保: 养老 8% + 医疗 2% + 失业 0.5% = 10.5%")
			xcolor.Success(config.Success, "公积金: 7%")
			xcolor.Success(config.Separator)

			ans := &config.TaxAnswer{}
			err := survey.Ask(config.TaxQuestion, ans)
			if err != nil {
				xcolor.Fail(config.Error, fmt.Sprintf("输入错误: %s", err.Error()))
				return
			}

			var (
				totalBase      float64 // 累计扣除基数
				totalSalary    float64 // 累计税前收入
				totalInsurance float64 // 累计社保
				totalFund      float64 // 累计公积金
				totalDeduction float64 // 累计附加扣除
				taxSalary      float64 // 累计计税收入
				lastTotalTax   float64 // 上月累计应纳税额
				totalTax       float64 // 累计应纳数额
				totalIncome    float64 // 累计税后收入
			)

			list := make([]*config.TaxIncome, 0)
			insurance := ans.Salary * ans.Insurance
			fund := ans.Salary * ans.Fund
			for month := 1; month <= 12; month++ {
				totalBase += TaxBase
				totalSalary += ans.Salary       // 累计税前收入
				totalInsurance += insurance     // 累计社保
				totalFund += fund               // 累计公积金
				totalDeduction += ans.Deduction // 累计附加扣除

				taxSalary = totalSalary - totalInsurance - totalFund - totalDeduction - totalBase
				taxRate, taxDeduction := internal.GetTaxRate(taxSalary)
				totalTax = taxSalary*taxRate - taxDeduction
				// 当月个税
				tax := totalTax - lastTotalTax
				lastTotalTax = totalTax
				// 当月税后
				income := ans.Salary - insurance - fund - tax
				totalIncome += income
				list = append(list, &config.TaxIncome{
					Month:       month,
					Salary:      xconvert.ToUint(ans.Salary),
					TotalSalary: xconvert.ToUint(totalSalary),
					Insurance:   fmt.Sprintf("%.0f (%.0f)", insurance, totalInsurance),
					Fund:        fmt.Sprintf("%.0f (%.0f)", fund, totalFund),
					Tax:         fmt.Sprintf("%.0f (%.0f)", tax, totalTax),
					Income:      xconvert.ToUint(income),
					TotalIncome: xconvert.ToUint(totalIncome),
				})
			}
			xtable.New(list).Style(xtable.Dashed).Border(true).Render()

			xcolor.Success(config.Separator)
		},
	}
)
