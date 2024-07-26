package xretry

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func request(num int) error {
	fmt.Println("request num", num)
	if num < 3 {
		return fmt.Errorf("invalid num")
	}
	return nil
}

var num = 1

func retry() bool {
	err := request(num)
	if err != nil {
		num++
	}
	return err == nil
}

func TestDo(t *testing.T) {
	assert.False(t, Do(retry, WithMax(1)))
}
