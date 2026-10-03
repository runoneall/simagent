package mcptools

import (
	"log"
	"simagent/alltools"
	"simagent/config"
	"simagent/onexit"

	"github.com/cloudwego/eino/components/tool"
)

var tools []tool.BaseTool

func init() {
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

	alltools.Add(tools...)
}
