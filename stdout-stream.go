package logger

import (
	"fmt"
)

type (
	StdOutStream struct {
		ILogStream
	}
)

func (s *StdOutStream) Write(msg LogMessage) {
	// Print the log message to the console
	fmt.Printf(
		"%s[%s] %s%s%s %6s %s[%s]%s %s%s\n", // Format string
		msg.Color,                           // Set colour
		msg.AppName,                         // Add the Prefix (app name)
		White,
		msg.Time, // Add the time (time is white)
		msg.Color,
		msg.Level, // Add the log level
		Yellow,
		msg.Name, // Add the log name (name of the logger is yellow)
		msg.Color,
		msg.Msg, // Add the message
		Reset,   // Reset colour
	)
}

func NewStdOutStream() *StdOutStream {
	return &StdOutStream{}
}
