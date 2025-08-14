package xopenai

import (
	"context"
	"fmt"
	"testing"
)

func TestRequestCompletion(t *testing.T) {
	resp, err := RequestCompletion(context.Background(), "你好")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
