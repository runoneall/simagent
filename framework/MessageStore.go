package framework

import (
	"sync"

	"github.com/barkimedes/go-deepcopy"
	"github.com/cloudwego/eino/schema"
)

const MessageStoreKey = "MessageStore"

type MessageStore struct {
	lock     sync.RWMutex
	messages []*schema.Message
}

func (ms *MessageStore) Set(messages []*schema.Message) {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	ms.messages = deepcopy.MustAnything(messages).([]*schema.Message)
}

func (ms *MessageStore) Get() []*schema.Message {
	ms.lock.RLock()
	defer ms.lock.RUnlock()
	return deepcopy.MustAnything(ms.messages).([]*schema.Message)
}

func (ms *MessageStore) Last() *schema.Message {
	ms.lock.RLock()
	defer ms.lock.RUnlock()

	length := len(ms.messages)
	if length == 0 {
		return nil
	}

	return deepcopy.MustAnything(ms.messages[length-1]).(*schema.Message)
}

func (ms *MessageStore) Append(message *schema.Message) {
	ms.lock.Lock()
	defer ms.lock.Unlock()
	ms.messages = append(ms.messages, deepcopy.MustAnything(message).(*schema.Message))
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		messages: []*schema.Message{},
	}
}
