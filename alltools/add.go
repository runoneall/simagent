package alltools

import (
	"sync"

	"github.com/cloudwego/eino/components/tool"
)

var (
	tools = []tool.BaseTool{}
	lock  sync.RWMutex
)

func Add(t ...tool.BaseTool) {
	lock.Lock()
	defer lock.Unlock()

	tools = append(tools, t...)
}
