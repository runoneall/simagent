package framework

import (
	"simagent/config"
	"simagent/threadmgr"

	"github.com/cloudwego/eino-ext/components/model/openai"
)

func NewChatModel() (*openai.ChatModel, error) {
	var (
		cfg                     = config.Get()
		MaxTokens           int = 16384
		MaxCompletionTokens     = MaxTokens
	)

	return openai.NewChatModel(threadmgr.Context, &openai.ChatModelConfig{
		BaseURL:             cfg.OpenAI.BaseURL,
		APIKey:              cfg.OpenAI.APIKey,
		Model:               cfg.OpenAI.Model,
		MaxTokens:           &MaxTokens,
		MaxCompletionTokens: &MaxCompletionTokens,
	})
}
