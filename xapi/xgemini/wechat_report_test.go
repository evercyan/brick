package xgemini

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/evercyan/brick/xfile"
)

func TestGenerateReport(t *testing.T) {
	ctx := context.Background()
	date := "2025-06-13"
	reporterDir := filepath.Join(xfile.GetHomeDir(), "Y1ker/Database/Chatlog/reporter")
	templatePath := filepath.Join(reporterDir, "群聊日报.txt")
	chatPath := filepath.Join(reporterDir, fmt.Sprintf("勇敢牛牛/勇敢牛牛_%s.txt", date))
	targetPath := filepath.Join(reporterDir, fmt.Sprintf("勇敢牛牛/勇敢牛牛_%s.html", date))
	err := GenerateReport(
		ctx, templatePath, chatPath, targetPath,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("targetPath:", targetPath)
}
