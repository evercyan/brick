package xutil

import (
	"fmt"
	"testing"
)

func TestCalculateColor(t *testing.T) {
	ranges := []float64{-10, -5, 0, 5, 10}
	colors := []string{"#196c2e", "#28a745", "#ffffff", "#dc3545", "#a71d2a"}
	values := []float64{-11, -10, -7, -5, -3, -1, 0, 1, 3, 5, 7, 10, 12}
	for _, value := range values {
		color, err := CalculateColor(value, ranges, colors)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Print(color + ",")
	}
}
