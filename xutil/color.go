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

// CalculateColor 计算颜色
func CalculateColor(value float64, ranges []float64, colors []string) (string, error) {
	if len(ranges) != len(colors) {
		return "", fmt.Errorf("颜色范围和选项数量不匹配")
	}
	if len(ranges) < 2 {
		return "", fmt.Errorf("无效的颜色范围")
	}
	factor, color1, color2 := math.MaxFloat64, "", ""
	if value < ranges[0] {
		factor = 0.0
		color1, color2 = colors[0], colors[1]
	} else if value > ranges[len(ranges)-1] {
		factor = 1.0
		color1, color2 = colors[len(ranges)-2], colors[len(ranges)-1]
	} else {
		match := -1
		for i := 0; i < len(ranges)-1; i++ {
			if value >= ranges[i] && value <= ranges[i+1] {
				match = i
				break
			}
		}
		if match == -1 {
			return "", fmt.Errorf("当前数值在给定范围中匹配到区间")
		}
		color1, color2 = colors[match], colors[match+1]
		vmin, vmax := ranges[match], ranges[match+1]
		if value == vmin {
			factor = 0.0
		} else {
			factor = (value - vmin) / (vmax - vmin)
		}
	}
	if factor == math.MaxFloat64 {
		return "", fmt.Errorf("计算颜色错误")
	}
	return InterpolateColor(factor, color1, color2), nil
}
