package config

type Config struct {
	OpenAI OpenAIConfig `json:"openai"`
	MCP    []MCPConfig  `json:"mcp"`
}

type OpenAIConfig struct {
	BaseURL string `json:"base_url"`
	APIKey  string `json:"api_key"`
	Model   string `json:"model"`
}

type MCPConfig struct {
	Type string `json:"type"`

	URL        string            `json:"url"`
	HTTPHeader map[string]string `json:"http_header"`

	Command string   `json:"command"`
	EnvVar  []string `json:"env_var"`
}

const (
	MCPTypeHTTP  = "http"
	MCPTypeStdio = "stdio"
)
