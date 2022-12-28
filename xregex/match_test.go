package xregex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchChinese(t *testing.T) {
	assert.Equal(t, []string{"你", "好呀"}, MatchChinese("h你e好呀llo"))
}
