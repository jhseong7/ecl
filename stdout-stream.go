package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
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
	NestJsStyle string = "nest"
	SpringStyle string = "spring"
)

func colourize(color string, msg string) string {
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
		s = fmt.Sprintf("%s%*s", s, n-len(s), "")
	}
	return s
}

func printNestStyleLog(msg LogMessage) {
	pid := os.Getpid()

	// If the Name is empty, then set it to the default value
	if msg.Name == "" {
		msg.Name = "default"
	}

	fmt.Printf(
		"%s %-7s - %s %s %s %s\n",                                    // Format string
		colourize(msg.Color, "["+msg.AppName+"]"),                    // Set colour
		colourize(msg.Color, padMinWidthRight(strconv.Itoa(pid), 6)), // Add the process id
		colourize(White, msg.Time.Format("01/02/2006, 3:04:05 PM")),  // Add the time (time is white)
		colourize(msg.Color, padMinWidthLeft(msg.Level, 6)),          // Add the log level
		colourize(Yellow, "["+msg.Name+"]"),                          // Add the log name (name of the logger is yellow)
		colourize(msg.Color, msg.Msg),                                // Add the message
	)
}

func printSpringStyleLog(msg LogMessage) {
	pid := os.Getpid()
	thread := "main" // Thread is always main

	timeStr := msg.Time.Format(time.RFC3339)

	// Split the date-time
	date := timeStr[:10]
	time := timeStr[11:]

	fmt.Printf(
		"%s %s %s --- %s %s %s\n",                           // <date-time>  <log level> <process id> --- [<thread>] <logger> : <message>
		colourize(White, date+" "+time),                     // Add the date-time (time is white)
		colourize(msg.Color, padMinWidthLeft(msg.Level, 6)), // Add the log level
		colourize(White, fmt.Sprintf("%d", pid)),            // Add the process id
		colourize(Yellow, "["+thread+"]"),                   // Add the thread
		colourize(Yellow, padMinWidthRight(msg.Name, 20)),   // Add the log name (name of the logger is yellow)
		colourize(msg.Color, msg.Msg),                       // Add the message
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
