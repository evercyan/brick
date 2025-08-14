package xopenai

import (
	"context"
	"fmt"

	"github.com/evercyan/brick/xutil"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// ...
var (
	Client openai.Client
)

// GetClient ...
func GetClient(ctx context.Context) (openai.Client, error) {
	if Client.Options != nil {
		return Client, nil
	}
	apiKey := xutil.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return Client, fmt.Errorf("OPENAI_API_KEY not found")
	}
	opts := []option.RequestOption{option.WithAPIKey(apiKey)}
	apiURL := xutil.Getenv("OPENAI_API_URL")
	if apiURL != "" {
		opts = append(opts, option.WithBaseURL(apiURL))
	}
	client := openai.NewClient(opts...)
	Client = client
	return Client, nil
}
