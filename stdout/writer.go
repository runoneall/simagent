package stdout

import (
	"context"
	"io"
	"os"
)

const DisableOutputKey = "DisableOutput"

type writer struct {
	out         io.Writer
	newLineFlag bool
	isFirstChar bool
}

var w = &writer{
	out:         os.Stdout,
	newLineFlag: false,
	isFirstChar: true,
}

func Writer(ctx context.Context) io.Writer {
	DisableOutput, ok := ctx.Value(DisableOutputKey).(bool)
	if ok && DisableOutput {
		return io.Discard
	}

	return w
}
