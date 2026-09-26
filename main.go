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

	"github.com/cloudwego/eino/schema"
)

func main() {
	threadmgr.Start(mainAgent)
	threadmgr.Wait()
}

func mainAgent() {
	ms := framework.NewMessageStore()
	ctx := context.WithValue(threadmgr.Context, "MessageStore", ms)

	runner, err := framework.NewRunner(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	iter := runner.Query(ctx, framework.UserPrompt())
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}

		if event.Err != nil {
			if errors.Is(event.Err, context.Canceled) {
				stdout.Logger.Println("INFO 用户中断")
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

	} else {
		iter := runner.Run(context.Background(), append(ms.Get(), schema.UserMessage("INTERRUPT")))
		for {
			event, ok := iter.Next()
			if !ok {
				break
			}

			if event.Err != nil {
				stdout.Logger.Println("ERROR", event.Err)
				break
			}

			if event.Output != nil && event.Output.MessageOutput != nil {
				stream := event.Output.MessageOutput.MessageStream

				if stream != nil {
					for {
						msg, err := stream.Recv()
						if errors.Is(err, io.EOF) {
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

		stdout.Logger.Println("INFO TinyAgent 已退出")
	}
}
