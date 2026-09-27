package main

import (
	"context"
	"errors"
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
	ms.Append(schema.UserMessage(framework.UserPrompt()))
	ctx := context.WithValue(threadmgr.Context, framework.MessageStoreKey, ms)

	if err := framework.Complete(ctx, ms); err != nil {
		if errors.Is(err, context.Canceled) {
			stdout.Logger.Println("INFO 用户取消任务")

		} else {
			stdout.Logger.Println("ERROR", err)
		}
	}

	stdout.Logger.Println("INFO Agent 已退出任务")

	if threadmgr.Context.Err() == nil {
		threadmgr.Start(mainAgent)

	} else if ms.Last().Role != schema.User {
		ms.Append(schema.UserMessage("INTERRUPT"))
		if err := framework.Complete(context.Background(), ms); err != nil {
			stdout.Logger.Println("ERROR", err)
		}

		stdout.Logger.Println("INFO TinyAgent 已退出")
	}
}
