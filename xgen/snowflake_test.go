package xgen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnowflake(t *testing.T) {
	NewSnowflake(1)
	NewSnowflake(10000)

	fmt.Println("SnowflakeID:", SnowflakeID())
	assert.NotEmpty(t, SnowflakeID())
	assert.NotEmpty(t, SnowflakeID())
	assert.NotEmpty(t, SnowflakeID())
}

// BenchmarkSnowflake-8   	 2196301	       548.0 ns/op	       0 B/op	       0 allocs/op
func BenchmarkSnowflake(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		SnowflakeID()
	}
}
