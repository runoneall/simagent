package msgstore

import "github.com/cloudwego/eino/schema"

func New() *MessageStore {
	return &MessageStore{
		messages: []*schema.Message{},
	}
}
