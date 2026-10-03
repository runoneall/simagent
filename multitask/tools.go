package multitask

import (
	"context"
	"fmt"
	"log"
	"simagent/framework"
	"simagent/msgstore"
	"simagent/stdout"
	"simagent/threadmgr"
	"strings"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/schema"
)

type (
	multitaskRunInput struct {
		Name  string `json:"name" jsonschema:"description=task name"`
		Input string `json:"input" jsonschema:"description=task input prompt"`
	}

	multitaskStatusInput struct {
		Name string `json:"name" jsonschema:"description=task name"`
	}

	multitaskResultInput struct {
		Name string `json:"name" jsonschema:"description=task name"`
	}

	multitaskDeleteInput struct {
		Name string `json:"name" jsonschema:"description=task name"`
	}

	multitaskListInput struct{}
)

const (
	multitaskRunDesc    = "create a task and run it in the background"
	multitaskStatusDesc = "get task status"
	multitaskResultDesc = "retrieve task results"
	multitaskDeleteDesc = "delete task"
	multitaskListDesc   = "list all tasks"
)

func Tools() []tool.BaseTool {
	multitaskRunTool, err := utils.InferTool(
		"multitask_run", multitaskRunDesc,
		func(_ context.Context, input *multitaskRunInput) (string, error) {
			lock.Lock()
			defer lock.Unlock()

			if _, exist := store[input.Name]; exist {
				return "", fmt.Errorf("task %s already exists", input.Name)
			}

			ms := msgstore.New()
			ms.Append(schema.UserMessage(input.Input))
			ctx := context.WithValue(
				context.WithValue(threadmgr.Context, stdout.DisableOutputKey, true),
				msgstore.MessageStoreKey, ms,
			)

			subagent := &task{ms: ms}
			threadmgr.Start(func() {
				err := framework.Complete(ctx, ms)

				lock.Lock()
				defer lock.Unlock()

				subagent.err = err
				subagent.exited = true
			})

			store[input.Name] = subagent
			return "success", nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	multitaskStatusTool, err := utils.InferTool(
		"multitask_status", multitaskStatusDesc,
		func(_ context.Context, input *multitaskStatusInput) (string, error) {
			lock.RLock()
			defer lock.RUnlock()

			subagent, exist := store[input.Name]
			if !exist {
				return "", fmt.Errorf("task %s does not exist", input.Name)
			}

			return fmt.Sprintf("error: %v, exited: %v", subagent.err, subagent.exited), nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	multitaskResultTool, err := utils.InferTool(
		"multitask_result", multitaskResultDesc,
		func(_ context.Context, input *multitaskResultInput) (string, error) {
			lock.RLock()
			defer lock.RUnlock()

			subagent, exist := store[input.Name]
			if !exist {
				return "", fmt.Errorf("task %s does not exist", input.Name)
			}

			if !subagent.exited {
				return "", fmt.Errorf("task %s has not exited", input.Name)
			}

			return subagent.ms.Last().Content, nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	multitaskDeleteTool, err := utils.InferTool(
		"multitask_delete", multitaskDeleteDesc,
		func(_ context.Context, input *multitaskDeleteInput) (string, error) {
			lock.Lock()
			defer lock.Unlock()

			subagent, exist := store[input.Name]
			if !exist {
				return "", fmt.Errorf("task %s does not exist", input.Name)
			}

			if !subagent.exited {
				return "", fmt.Errorf("task %s has not exited", input.Name)
			}

			delete(store, input.Name)
			return "success", nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	multitaskListTool, err := utils.InferTool(
		"multitask_list", multitaskListDesc,
		func(_ context.Context, input *multitaskListInput) (string, error) {
			lock.RLock()
			defer lock.RUnlock()

			tasks := []string{}
			for name, subagent := range store {
				tasks = append(tasks, fmt.Sprintf("task %s: {error: %v, exited: %v}", name, subagent.err, subagent.exited))
			}

			return strings.Join(tasks, "\n"), nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	return []tool.BaseTool{
		multitaskRunTool,
		multitaskStatusTool,
		multitaskResultTool,
		multitaskDeleteTool,
		multitaskListTool,
	}
}
