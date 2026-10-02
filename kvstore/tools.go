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
		Key string `json:"key" jsonschema:"description=The exact key to retrieve from the store"`
	}

	kvSetInput struct {
		Key   string `json:"key" jsonschema:"description=The exact key to create or overwrite"`
		Value string `json:"value" jsonschema:"description=The value to store for the key"`
	}

	kvDeleteInput struct {
		Key string `json:"key" jsonschema:"description=The exact key to delete from the store"`
	}

	kvListInput struct{}
)

const (
	kvGetDesc    = "Read a value from the key-value store by key. This is a read-only operation and may be executed concurrently with other read-only tools such as kv_list. Use it when you already know the key and need its value. Returns the stored value as a string."
	kvSetDesc    = "Write a key-value pair to the key-value store. This is a write operation: call it by itself, never in parallel with any other tool call, and wait for the result before continuing. Use it to create or overwrite a key. The store has limited space; delete unused keys first if needed. Returns \"success\" when the write succeeds."
	kvDeleteDesc = "Delete a key from the key-value store. This is a write operation: call it by itself, never in parallel with any other tool call, and wait for the result before continuing. Use it to remove keys that are no longer needed, especially before kv_set when space is limited. Returns \"success\" when the delete succeeds."
	kvListDesc   = "List all keys currently present in the key-value store. This is a read-only operation and may be executed concurrently with other read-only tools such as kv_get. Use it to discover available keys before reading, updating, or deleting them. Returns only the keys, not their values."
)

func Tools() []tool.BaseTool {
	kvGetTool, err := utils.InferTool(
		"kv_get", kvGetDesc,
		func(_ context.Context, input *kvGetInput) (string, error) {
			return kvGet(input.Key)
		},
	)

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
