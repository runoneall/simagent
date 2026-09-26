package framework

import (
	"context"

	"github.com/cloudwego/eino/adk"
)

func NewRunner(ctx context.Context) (*adk.Runner, error) {
	agent, err := NewAgent(ctx)
	if err != nil {
		return nil, err
	}

	return adk.NewRunner(ctx, adk.RunnerConfig{
		Agent:           agent,
		EnableStreaming: true,
	}), nil
}
