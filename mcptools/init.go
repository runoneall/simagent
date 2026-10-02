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

	for _, cfg := range config.Get().MCP {
		client, err := mcpclient(cfg)
		if err != nil {
			log.Fatalln(err)
		}

		onexit.Do(func() {
			client.Close()
		})

		mcpTools, err := getmcptools(client)
		if err != nil {
			log.Fatalln(err)
		}

		tools = append(tools, mcpTools...)
	}
}
