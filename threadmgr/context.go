package threadmgr

import "context"

var Context, cancel = context.WithCancel(context.Background())
