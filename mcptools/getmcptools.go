package mcptools

import (
	"simagent/threadmgr"

	mcpp "github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/components/tool"
	mcpc "github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func getmcptools(client *mcpc.Client) ([]tool.BaseTool, error) {
	if err := client.Start(threadmgr.Context); err != nil {
		return nil, err
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "simagent-mcp-client",
		Version: "1.0.0",
	}

	if _, err := client.Initialize(threadmgr.Context, initRequest); err != nil {
		return nil, err
	}

	return mcpp.GetTools(threadmgr.Context, &mcpp.Config{Cli: client})
}
