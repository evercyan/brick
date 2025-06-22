package xgemini

import (
	"context"
	"fmt"
	"os"
	"strings"

	"google.golang.org/genai"
)

// https://ai.google.dev/gemini-api/docs/document-processing?hl=zh-cn&lang=go

// RequestFile ...
func RequestFile(ctx context.Context, sourcePath string) (string, error) {
	client, err := GetClient(ctx)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(sourcePath, ".txt") {
		return "", fmt.Errorf("输入文件必须是 txt 文件")
	}
	b, err := os.ReadFile(sourcePath)
	if err != nil {
		return "", err
	}
	parts := []*genai.Part{{
		InlineData: &genai.Blob{
			MIMEType: "text/plain",
			Data:     b,
		},
	}}
	contents := []*genai.Content{genai.NewContentFromParts(parts, genai.RoleUser)}
	result, err := client.Models.GenerateContent(ctx, ModelFlashPreview, contents, nil)
	if err != nil {
		return "", err
	}
	return result.Text(), nil
}
