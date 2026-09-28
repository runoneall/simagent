package mcptools

import (
	"fmt"
	"simagent/config"

	"github.com/arkady-emelyanov/go-shellparse"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/client/transport"
)

func newClient(cfg config.MCPConfig) (*client.Client, error) {
	switch cfg.Type {

	case config.MCPTypeHTTP:
		return client.NewStreamableHttpClient(cfg.URL, transport.WithHTTPHeaders(cfg.HTTPHeader))

	case config.MCPTypeStdio:
		bin, args, err := shellparse.Command(cfg.Command)
		if err != nil {
			return nil, err
		}

		return client.NewStdioMCPClient(bin, cfg.EnvVar, args...)

	}

	return nil, fmt.Errorf("未知 MCP 类型: %s", cfg.Type)
}
