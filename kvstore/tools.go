package kvstore

import (
	"context"
	"log"
	"strings"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type (
	kvGetInput struct {
		Key string `json:"key" jsonschema:"description=key"`
	}

	kvSetInput struct {
		Key   string `json:"key" jsonschema:"description=key"`
		Value string `json:"value" jsonschema:"description=value"`
	}

	kvDeleteInput struct {
		Key string `json:"key" jsonschema:"description=key"`
	}

	kvListInput struct{}
)

const (
	kvGetDesc    = "read the value from the KV store using the key"
	kvSetDesc    = "set the specified key in the KV store to the specified value"
	kvDeleteDesc = "delete the specified key from the KV store"
	kvListDesc   = "list all keys in the KV store"
)

func Tools() []tool.BaseTool {
	kvGetTool, err := utils.InferTool(
		"kv_get", kvGetDesc,
		func(_ context.Context, input *kvGetInput) (string, error) {
			return kvGet(input.Key)
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	kvSetTool, err := utils.InferTool(
		"kv_set", kvSetDesc,
		func(_ context.Context, input *kvSetInput) (string, error) {
			err := kvSet(input.Key, input.Value)
			if err != nil {
				return "", err
			}

			return "success", nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	kvDeleteTool, err := utils.InferTool(
		"kv_delete", kvDeleteDesc,
		func(_ context.Context, input *kvDeleteInput) (string, error) {
			err := kvDelete(input.Key)
			if err != nil {
				return "", err
			}

			return "success", nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	kvListTool, err := utils.InferTool(
		"kv_list", kvListDesc,
		func(_ context.Context, input *kvListInput) (string, error) {
			return strings.Join(kvList(), ","), nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	return []tool.BaseTool{kvGetTool, kvSetTool, kvDeleteTool, kvListTool}
}
