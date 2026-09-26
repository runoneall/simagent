package framework

import (
	"os"
	"strings"
)

const emptyUserPrompt = "你是谁？你可以做什么？当前时间是什么？"

func UserPrompt() string {
	content, err := os.ReadFile("agent.md")
	if err != nil {
		return emptyUserPrompt
	}

	prompt := string(content)
	if strings.TrimSpace(prompt) == "" {
		return emptyUserPrompt
	}

	return prompt
}
