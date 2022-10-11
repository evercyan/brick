package internal

import (
	"fmt"
)

// FormatPrice 价钱格式化
func FormatPrice(v float64) string {
	return fmt.Sprintf("%.2f", v)
}

// GetTaxRate 计算个税比率和速算金额
// 1, 不超过 36,000 元的部分, 税率: 3%, 速算扣除: 0
// 2, 超过 36,000 元至 144,000 元的部分, 税率: 10%, 速算扣除: 2520
// 3, 超过 144,000 元至 300,000 元的部分, 税率: 20%, 速算扣除: 16920
// 4, 超过 300,000 元至 420,000 元的部分, 税率: 25%, 速算扣除: 31920
// 5, 超过 420,000 元至 660,000 元的部分, 税率: 30%, 速算扣除: 52920
// 6, 超过 660,000 元至 960,000 元的部分, 税率: 35%, 速算扣除: 85920
// 7, 超过 960,000 元的部分, 税率: 45%, 速算扣除: 181920
func GetTaxRate(taxIncome float64) (float64, float64) {
	if taxIncome < 36000 {
		return 0.03, 0
	} else if taxIncome < 144000 {
		return 0.1, 2520
	} else if taxIncome < 300000 {
		return 0.2, 16920
	} else if taxIncome < 420000 {
		return 0.25, 31920
	} else if taxIncome < 660000 {
		return 0.3, 52920
	} else if taxIncome < 960000 {
		return 0.35, 85920
	}
	return 0.45, 181920
}
