package framework

import (
	_ "embed"
	"fmt"
	"math"
	"simagent/threadmgr"
	"time"

	"github.com/cloudwego/eino/adk"
	"github.com/valyala/fasttemplate"
)

//go:embed system.md
var systemPrompt string

func NewAgent() (*adk.ChatModelAgent, error) {
	chatModel, err := NewChatModel()
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

	return adk.NewChatModelAgent(threadmgr.Context, &adk.ChatModelAgentConfig{
		Model:         chatModel,
		MaxIterations: math.MaxInt,
		Instruction: fasttemplate.New(systemPrompt, "{{", "}}").ExecuteString(map[string]any{
			"CURRENT_TIME": currentTime,
			"TIME_ZONE":    timeZone,
		}),
	})
}
