package msgstore

import (
	"github.com/barkimedes/go-deepcopy"
	"github.com/cloudwego/eino/schema"
)

func (ms *MessageStore) Append(message *schema.Message) {
	ms.lock.Lock()
	defer ms.lock.Unlock()

	ms.messages = append(ms.messages, deepcopy.MustAnything(message).(*schema.Message))
}
