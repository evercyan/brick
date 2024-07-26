package xgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	assert.LessOrEqual(t, 1, RandInt(1, 10))
	assert.GreaterOrEqual(t, 10, RandInt(1, 10))

	assert.LessOrEqual(t, -10, RandInt(-10, 10))
	assert.GreaterOrEqual(t, 10, RandInt(-10, 10))
}

func TestRange(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Range(2, 1))
	assert.Equal(t, []int{1, 2}, Range(1, 2))
}

func TestRandString(t *testing.T) {
	assert.Equal(t, 6, len(RandStr(6)))
	assert.NotEqual(t, RandStr(6), RandStr(6))
}

// BenchmarkRandInt-8    51310586    24.29 ns/op    0 B/op    0 allocs/op
func BenchmarkRandInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandInt(10, 99)
	}
}

// BenchmarkRandInt-8    51310586    24.29 ns/op    0 B/op    0 allocs/op
func BenchmarkRandStr(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RandStr(10)
	}
}
