package xgemini

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/evercyan/brick/xfile"
)

func TestRequestFile(t *testing.T) {
	ctx := context.Background()
	sourcePath := filepath.Join(xfile.GetHomeDir(), "Y1ker/AI/交易复盘/交易复盘_0619.txt")
	resp, err := RequestFile(ctx, sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
