package threadmgr

import (
	"os"
	"os/signal"
)

func Wait() {
	sigInterrupt := make(chan os.Signal, 1)
	signal.Notify(sigInterrupt, os.Interrupt)
	<-sigInterrupt

	cancel()
	wg.Wait()
}
