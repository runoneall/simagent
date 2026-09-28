package framework

import (
	"context"
	_ "embed"
	"fmt"
	"math"
	"simagent/kvstore"
	"simagent/stdout"
	"time"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/compose"
	"github.com/valyala/fasttemplate"
)

//go:embed system.md
var systemPrompt string

func NewAgent(ctx context.Context) (*adk.ChatModelAgent, error) {
	chatModel, err := NewChatModel(ctx)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	currentTime := now.Format("2006-01-02 15:04:05")

	zoneName, offset := now.Zone()
	timeZone := fmt.Sprintf("%s (UTC%s%d)", zoneName, func() string {
		if offset >= 0 {
			return "+"
		}

		return ""
	}(), offset/3600)

	saveState := func(ctx context.Context, state *adk.ChatModelAgentState) error {
		ms, ok := ctx.Value(MessageStoreKey).(*MessageStore)
		if ok {
			ms.Set(state.Messages)
		}

		return nil
	}

	kvTools, err := kvstore.Tools()
	if err != nil {
		return nil, err
	}

	return adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Model:         chatModel,
		MaxIterations: math.MaxInt,

		Instruction: fasttemplate.New(systemPrompt, "{{", "}}").ExecuteString(map[string]any{
			"CURRENT_TIME": currentTime,
			"TIME_ZONE":    timeZone,
		}),

		Middlewares: []adk.AgentMiddleware{
			{
				BeforeChatModel: saveState,
				AfterChatModel:  saveState,
			},
		},

		ToolsConfig: adk.ToolsConfig{
			ToolsNodeConfig: compose.ToolsNodeConfig{
				Tools: kvTools,
				ToolCallMiddlewares: []compose.ToolMiddleware{
					{
						Invokable: func(next compose.InvokableToolEndpoint) compose.InvokableToolEndpoint {
							return func(ctx context.Context, input *compose.ToolInput) (*compose.ToolOutput, error) {
								stdout.Logger.Printf("INFO 工具调用 tool=%s arguments=%s\n", input.Name, input.Arguments)

								output, err := next(ctx, input)
								if err != nil {
									stdout.Logger.Println("ERROR", err)

									return &compose.ToolOutput{
										Result: fmt.Sprintf("Tool %s execution failed with error: %v. Please correct your input or handle this fallback.", input.Name, err),
									}, nil
								}

								return output, nil
							}
						},
					},
				},
			},
		},
	})
}
