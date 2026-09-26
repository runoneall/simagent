package framework

import (
	"simagent/threadmgr"

	"github.com/cloudwego/eino/adk"
)

func NewRunner() (*adk.Runner, error) {
	agent, err := NewAgent()
	if err != nil {
		return nil, err
	}

	return adk.NewRunner(threadmgr.Context, adk.RunnerConfig{
		Agent:           agent,
		EnableStreaming: true,
	}), nil
}
