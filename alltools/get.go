package alltools

import "github.com/cloudwego/eino/components/tool"

func Get() []tool.BaseTool {
	lock.RLock()
	defer lock.RUnlock()

	return tools
}
