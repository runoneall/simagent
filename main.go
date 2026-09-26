package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"simagent/framework"
	"simagent/stdout"
	"simagent/threadmgr"
	"time"
)

func main() {
	threadmgr.Start(mainAgent)
	threadmgr.Wait()
}

func mainAgent() {
	runner, err := framework.NewRunner()
	if err != nil {
		log.Fatalln(err)
	}

	iter := runner.Query(threadmgr.Context, framework.UserPrompt())
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}

		if event.Err != nil {
			if errors.Is(event.Err, context.Canceled) {
				time.Sleep(5 * time.Second)
				break
			}

			stdout.Logger.Println("ERROR", event.Err)
			break
		}

		if event.Output != nil && event.Output.MessageOutput != nil {
			stream := event.Output.MessageOutput.MessageStream

			if stream != nil {
				for {
					msg, err := stream.Recv()
					if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
						break
					}

					if err != nil {
						stdout.Logger.Println("ERROR", err)
						break
					}

					if msg != nil {
						fmt.Fprint(stdout.Writer, msg.Content)
					}
				}
			}
		}
	}

	stdout.Logger.Println("INFO Agent 已退出任务")

	if threadmgr.Context.Err() == nil {
		threadmgr.Start(mainAgent)
	}
}
