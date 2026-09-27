package framework

import (
	"context"
	"errors"
	"fmt"
	"io"
	"simagent/stdout"
)

func Complete(ctx context.Context, ms *MessageStore) error {
	runner, err := NewRunner(ctx)
	if err != nil {
		return err
	}

	iter := runner.Run(ctx, ms.Get())
	for {
		event, ok := iter.Next()
		if !ok {
			break
		}

		if event.Err != nil {
			return event.Err
		}

		if event.Output != nil && event.Output.MessageOutput != nil {
			stream := event.Output.MessageOutput.MessageStream

			if stream != nil {
				for {
					msg, err := stream.Recv()
					if errors.Is(err, io.EOF) {
						break
					}

					if err != nil {
						return err
					}

					if msg != nil {
						fmt.Fprint(stdout.Writer, msg.Content)
					}
				}
			}
		}
	}

	return nil
}
