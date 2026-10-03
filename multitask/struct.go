package multitask

import "simagent/msgstore"

type task struct {
	ms     *msgstore.MessageStore
	err    error
	exited bool
}
