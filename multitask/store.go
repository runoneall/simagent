package multitask

import "sync"

var (
	store = map[string]*task{}
	lock  sync.RWMutex
)
