package msgstore

import (
	"github.com/barkimedes/go-deepcopy"
	"github.com/cloudwego/eino/schema"
)

func (ms *MessageStore) Set(messages []*schema.Message) {
	ms.lock.Lock()
	defer ms.lock.Unlock()

	ms.messages = deepcopy.MustAnything(messages).([]*schema.Message)
}
