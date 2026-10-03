package msgstore

import (
	"github.com/barkimedes/go-deepcopy"
	"github.com/cloudwego/eino/schema"
)

func (ms *MessageStore) Get() []*schema.Message {
	ms.lock.RLock()
	defer ms.lock.RUnlock()

	return deepcopy.MustAnything(ms.messages).([]*schema.Message)
}
