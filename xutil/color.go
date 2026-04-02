package xutil

import (
	"fmt"
	"math"

	"github.com/evercyan/brick/xconvert"
)

// InterpolateColor 在两个颜色之间进行线性插值
// factor 为插值因子(0.0 到 1.0), 0 返回 color1, 1 返回 color2
func InterpolateColor(factor float64, color1, color2 string) string {
	r1, g1, b1 := xconvert.Hex2RGB(color1)
	r2, g2, b2 := xconvert.Hex2RGB(color2)
	r := float64(r1) + factor*(float64(r2)-float64(r1))
	g := float64(g1) + factor*(float64(g2)-float64(g1))
	b := float64(b1) + factor*(float64(b2)-float64(b1))
	r = math.Max(0, math.Min(255, r))
	g = math.Max(0, math.Min(255, g))
	b = math.Max(0, math.Min(255, b))
	return "#" + xconvert.RGB2Hex(int(r), int(g), int(b))
}

// CalculateColor 根据数值和颜色配置计算颜色
// thresholds: 数值阈值点，如 []float64{0, 50, 100}
// colors: 对应阈值点的颜色，如 []string{"#00FF00", "#FFFF00", "#FF0000"}
// 当数值落在 thresholds[i] 到 thresholds[i+1] 之间时，在 colors[i] 到 colors[i+1] 之间线性插值
func CalculateColor(value float64, thresholds []float64, colors []string) (string, error) {
	if len(thresholds) != len(colors) {
		return "", fmt.Errorf("thresholds 和 colors 长度必须相等")
	}
	if len(thresholds) < 2 {
		return "", fmt.Errorf("thresholds 和 colors 至少需要 2 个元素")
	}
	for i := 1; i < len(thresholds); i++ {
		if thresholds[i] <= thresholds[i-1] {
			return "", fmt.Errorf("thresholds 必须严格递增")
		}
	}
	if value <= thresholds[0] {
		return colors[0], nil
	}
	if value >= thresholds[len(thresholds)-1] {
		return colors[len(thresholds)-1], nil
	}
	for i := 0; i < len(thresholds)-1; i++ {
		if value >= thresholds[i] && value <= thresholds[i+1] {
			factor := (value - thresholds[i]) / (thresholds[i+1] - thresholds[i])
			return InterpolateColor(factor, colors[i], colors[i+1]), nil
		}
	}
	return "", fmt.Errorf("未找到匹配的区间")
}
