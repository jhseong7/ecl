package stream

import "github.com/jhseong7/ecl/message"

type (
	ILogStream interface {
		Write(msg message.LogMessage)

		// TODO: add a flush method so the logger can handle any pending messages before the program exits
		// Flush()
	}
)
