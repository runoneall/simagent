package stdout

import (
	"context"
	"log"
)

func Logger(ctx context.Context) *log.Logger {
	return log.New(Writer(ctx), "\n", log.LstdFlags)
}
