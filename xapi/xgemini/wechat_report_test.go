package xgemini

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/evercyan/brick/xfile"
)

func TestGenerateWechatDailyReport(t *testing.T) {
	ctx := context.Background()
	chatlogDir := filepath.Join(xfile.GetHomeDir(), "Y1ker/AI/chatlog")
	// 微信聊天记录提示词
	// 勇敢牛牛_20250611
	templatePath := filepath.Join(chatlogDir, "微信聊天记录提示词.txt")
	chatPath := filepath.Join(chatlogDir, "勇敢牛牛_20250611.txt")
	targetPath := filepath.Join(chatlogDir, "勇敢牛牛_202506111111.html")
	err := GenerateWechatReport(
		ctx, templatePath, chatPath, targetPath,
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("targetPath:", targetPath)
}
