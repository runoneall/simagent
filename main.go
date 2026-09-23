package main

import (
	"fmt"
	"simagent/threadmgr"
	"time"
)

func main() {
	threadmgr.Start(mainAgent)
	threadmgr.Wait()
}

func mainAgent() {
	time.Sleep(5 * time.Second)
	fmt.Println(1)

	if threadmgr.Context.Err() == nil {
		threadmgr.Start(mainAgent)
	}
}
