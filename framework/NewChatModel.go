package framework

import (
	"context"
	"simagent/config"

	"github.com/cloudwego/eino-ext/components/model/openai"
)

func NewChatModel(ctx context.Context) (*openai.ChatModel, error) {
	var (
		cfg                     = config.Get()
		MaxTokens           int = 16384
		MaxCompletionTokens     = MaxTokens
	)

	return openai.NewChatModel(ctx, &openai.ChatModelConfig{
		BaseURL:             cfg.OpenAI.BaseURL,
		APIKey:              cfg.OpenAI.APIKey,
		Model:               cfg.OpenAI.Model,
		MaxTokens:           &MaxTokens,
		MaxCompletionTokens: &MaxCompletionTokens,
	})
}
