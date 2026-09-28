package onexit

import (
	"simagent/threadmgr"
	"sync"
)

var (
	jobs = []func(){}
	lock sync.RWMutex
)

func init() {
	threadmgr.Start(func() {
		<-threadmgr.Context.Done()

		lock.RLock()
		defer lock.RUnlock()

		for _, job := range jobs {
			job()
		}
	})
}
