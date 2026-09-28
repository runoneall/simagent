package mcptools

import (
	"log"
	"simagent/config"
	"simagent/onexit"
	"sync"

	"github.com/cloudwego/eino/components/tool"
)

var (
	tools []tool.BaseTool
	lock  sync.RWMutex
)

func init() {
	lock.Lock()
	defer lock.Unlock()

	cfg := config.Get()
	for _, mcpCfg := range cfg.MCP {
		cli, err := newClient(mcpCfg)
		if err != nil {
			log.Fatalln(err)
		}

		onexit.Do(func() {
			cli.Close()
		})

		mcpTool, err := getMCPTool(cli)
		if err != nil {
			log.Fatalln(err)
		}

		tools = append(tools, mcpTool...)
	}
}
