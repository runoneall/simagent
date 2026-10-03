package multitask

import (
	"simagent/alltools"
	"sync"
)

var (
	store = map[string]*task{}
	lock  sync.RWMutex
)

func init() {
	alltools.Add(tools()...)
}
