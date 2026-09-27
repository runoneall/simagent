package framework

import (
	"sync"

	"github.com/cloudwego/eino/schema"
)

const MessageStoreKey = "MessageStore"

type MessageStore struct {
	lock     sync.RWMutex
	messages []*schema.Message
}

func (ms *MessageStore) Set(messages []*schema.Message) {
	ms.lock.Lock()
	ms.messages = messages
	ms.lock.Unlock()
}

func (ms *MessageStore) Get() []*schema.Message {
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	return ms.messages
}

func (ms *MessageStore) Last() *schema.Message {
	ms.lock.RLock()
	defer ms.lock.RUnlock()

	length := len(ms.messages)
	if length == 0 {
		return nil
	}

	return ms.messages[length-1]
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		messages: []*schema.Message{},
	}
}
