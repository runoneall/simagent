package msgstore

import (
	"sync"

	"github.com/cloudwego/eino/schema"
)

const MessageStoreKey = "MessageStore"

type MessageStore struct {
	lock     sync.RWMutex
	messages []*schema.Message
}
