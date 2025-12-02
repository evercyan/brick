package xutil

import (
	"fmt"
	"testing"
)

func TestCalculateColor(t *testing.T) {
	ranges := []float64{-10, 0, 10}
	colors := []string{"#00B050", "#FFFFFF", "#C00000"}
	values := []float64{-15, -10, -5, -1, 0, 1, 5, 10, 15}
	for _, value := range values {
		color, err := CalculateColor(value, ranges, colors)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(value, color)
	}
}
