package azure

import (
	"context"

	"github.com/starmvp/langchaingo/llms"
)

type LLM struct {
	endpoint string
	apiKey   string
}

func New(endpoint, apiKey string) (*LLM, error) {
	return &LLM{
		endpoint: endpoint,
		apiKey:   apiKey,
	}, nil
}

func (a *LLM) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	// Here you would call the Azure LLM API
	response := &llms.ContentResponse{}
	return response, nil
}

func (a *LLM) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return a.GenerateFromSinglePrompt(ctx, prompt, options...)
}

func (a *LLM) GenerateFromSinglePrompt(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	return llms.GenerateFromSinglePrompt(ctx, a, prompt, options...)
}
