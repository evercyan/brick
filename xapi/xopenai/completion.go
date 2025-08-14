package xopenai

import (
	"context"

	"github.com/openai/openai-go"
)

// RequestCompletion
func RequestCompletion(ctx context.Context, message string) (string, error) {
	client, err := GetClient(ctx)
	if err != nil {
		return "", err
	}
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
		Model: "deepseek-ai/DeepSeek-R1",
	})
	if err != nil {
		return "", err
	}
	return chatCompletion.Choices[0].Message.Content, nil
}
