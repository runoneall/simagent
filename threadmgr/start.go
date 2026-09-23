package threadmgr

import "sync"

var wg sync.WaitGroup

func Start(job func()) {
	wg.Go(job)
}
