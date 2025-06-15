package xgemini

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/evercyan/brick/xfile"
	"google.golang.org/genai"
)

// https://ai.google.dev/gemini-api/docs/document-processing?hl=zh-cn&lang=go

// GenerateWechatReport ...
func GenerateWechatReport(ctx context.Context, templatePath, chatPath, targetPath string) error {
	client, err := GetClient(ctx)
	if err != nil {
		return err
	}
	parts := make([]*genai.Part, 0)
	for _, v := range []string{templatePath, chatPath} {
		if !strings.HasSuffix(v, ".txt") {
			return fmt.Errorf("输入文件必须是 txt 文件")
		}
		b, err := os.ReadFile(v)
		if err != nil {
			return err
		}
		parts = append(parts, &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: "text/plain",
				Data:     b,
			},
		})
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}
	result, err := client.Models.GenerateContent(ctx, ModelGeminiPro, contents, nil)
	if err != nil {
		return err
	}
	return xfile.Write(targetPath, result.Text())
}
