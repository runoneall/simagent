package msgstore

import (
	"github.com/barkimedes/go-deepcopy"
	"github.com/cloudwego/eino/schema"
)

func (ms *MessageStore) Last() *schema.Message {
	ms.lock.RLock()
	defer ms.lock.RUnlock()

	length := len(ms.messages)
	if length == 0 {
		return nil
	}

	return deepcopy.MustAnything(ms.messages[length-1]).(*schema.Message)
}
