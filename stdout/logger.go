package stdout

import "log"

var Logger = log.New(Writer, "\n", log.LstdFlags)
