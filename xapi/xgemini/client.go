package xgemini

import (
	"context"
	"fmt"

	"github.com/evercyan/brick/xutil"
	"google.golang.org/genai"
)

// ...
var (
	Client *genai.Client
)

// GetClient ...
func GetClient(ctx context.Context) (*genai.Client, error) {
	if Client != nil {
		return Client, nil
	}
	apiKey := xutil.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY not found")
	}
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	Client = client
	return Client, nil
}
