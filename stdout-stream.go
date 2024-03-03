package logger

import (
	"fmt"
	"os"
)

type (
	StdOutStream struct {
		ILogStream
		logStyle string
	}

	StdOutStreamOption struct {
		LogStyle string
	}
)

const (
	NestJsStyle = "nest"
	SpringStyle = "spring"
)

func colorize(color string, msg string) string {
	return fmt.Sprintf("%s%s%s", color, msg, Reset)
}

// Pad the min width of a string to the front
func padMinWidthLeft(s string, n int) string {
	if len(s) < n {
		s = fmt.Sprintf("%*s%s", n-len(s), "", s)
	}
	return s
}

// Pad the min width of a string to the back
func padMinWidthRight(s string, n int) string {
	if len(s) < n {
		s = fmt.Sprintf("%-*s%s", n-len(s), "", s)
	}
	return s
}

func printNestStyleLog(msg LogMessage) {
	fmt.Printf(
		"%s %s %s %s %s\n",                                 // Format string
		colorize(msg.Color, "["+msg.AppName+"]"),           // Set colour
		colorize(White, msg.Time),                          // Add the time (time is white)
		colorize(msg.Color, padMinWidthLeft(msg.Level, 6)), // Add the log level
		colorize(Yellow, "["+msg.Name+"]"),                 // Add the log name (name of the logger is yellow)
		colorize(msg.Color, msg.Msg),                       // Add the message
	)
}

func printSpringStyleLog(msg LogMessage) {
	pid := os.Getpid()
	thread := "main" // Thread is always main

	// Split the date-time
	date := msg.Time[:10]
	time := msg.Time[11:]

	fmt.Printf(
		"%s %s %s --- %s %s %s\n",                          // <date-time>  <log level> <process id> --- [<thread>] <logger> : <message>
		colorize(White, date+" "+time),                     // Add the date-time (time is white)
		colorize(msg.Color, padMinWidthLeft(msg.Level, 6)), // Add the log level
		colorize(White, fmt.Sprintf("%d", pid)),            // Add the process id
		colorize(Yellow, "["+thread+"]"),                   // Add the thread
		colorize(Yellow, padMinWidthRight(msg.Name, 20)),   // Add the log name (name of the logger is yellow)
		colorize(msg.Color, msg.Msg),                       // Add the message
	)
}

func (s *StdOutStream) Write(msg LogMessage) {
	// Print the log message to the console
	switch s.logStyle {
	case NestJsStyle:
		printNestStyleLog(msg)
	case SpringStyle:
		printSpringStyleLog(msg)
	default:
		printNestStyleLog(msg)
	}
}

func NewStdOutStream(options ...StdOutStreamOption) *StdOutStream {
	// if len > 1, then it's an error
	if len(options) > 1 {
		panic("NewStdOutStream: Too many options")
	}

	// When an explicit log style is given
	if len(options) == 1 {
		return &StdOutStream{
			logStyle: options[0].LogStyle,
		}
	}

	// if the option is given via environment variable
	logStyle := os.Getenv("LOG_STYLE")
	if logStyle != "" {
		return &StdOutStream{
			logStyle: logStyle,
		}
	}

	return &StdOutStream{}
}
