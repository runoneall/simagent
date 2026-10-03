package framework

import (
	"context"
	_ "embed"
	"fmt"
	"math"
	"simagent/kvstore"
	"simagent/mcptools"
	"simagent/msgstore"
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
		ms, ok := ctx.Value(msgstore.MessageStoreKey).(*msgstore.MessageStore)
		if ok {
			ms.Set(state.Messages)
		}

		return nil
	}

	logger := stdout.Logger(ctx)
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
				Tools: append(kvstore.Tools(), mcptools.Tools()...),
				ToolCallMiddlewares: []compose.ToolMiddleware{
					{
						Invokable: func(next compose.InvokableToolEndpoint) compose.InvokableToolEndpoint {
							return func(ctx context.Context, input *compose.ToolInput) (*compose.ToolOutput, error) {
								logger.Printf("INFO 工具调用 tool=%s arguments=%s\n", input.Name, input.Arguments)

								output, err := next(ctx, input)
								if err != nil {
									logger.Println("ERROR", err)

									return &compose.ToolOutput{
										Result: fmt.Sprintf("[TOOL ERROR] tool '%s' failed: %v. please correct your parameters and retry, or explain the failure and proceed with a fallback response.", input.Name, err),
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
