package main

import (
	"context"
	"errors"
	"simagent/framework"
	"simagent/msgstore"
	"simagent/stdout"
	"simagent/threadmgr"

	"github.com/cloudwego/eino/schema"
)

func main() {
	threadmgr.Start(mainAgent)
	threadmgr.Wait()
}

func mainAgent() {
	ms := msgstore.New()
	ms.Append(schema.UserMessage(framework.UserPrompt()))
	ctx := context.WithValue(threadmgr.Context, msgstore.MessageStoreKey, ms)
	logger := stdout.Logger(ctx)

	if err := framework.Complete(ctx, ms); err != nil {
		if errors.Is(err, context.Canceled) {
			logger.Println("INFO 用户取消任务")

		} else {
			logger.Println("ERROR", err)
		}
	}

	logger.Println("INFO Agent 已退出任务")

	if threadmgr.Context.Err() == nil {
		threadmgr.Start(mainAgent)

	} else if ms.Last().Role != schema.User {
		ms.Append(schema.UserMessage("[SYSTEM MESSAGE]: INTERRUPT"))
		if err := framework.Complete(context.Background(), ms); err != nil {
			logger.Println("ERROR", err)
		}

		logger.Println("INFO SimAgent 已退出")
	}
}
