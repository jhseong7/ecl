package style

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jhseong7/ecl/message"
)

type (
	LogStyle string
)

const (
	DefaultStyle LogStyle = "DEFAULT"
	NestJsStyle  LogStyle = "NESTJS"
	SpringStyle  LogStyle = "SPRING"
)

func colourize(color string, msg string) string {
	return fmt.Sprintf("%s%s%s", color, msg, Reset)
}

func bold(msg string) string {
	return fmt.Sprintf("%s%s%s", Bold, msg, Reset)
}

func italic(msg string) string {
	return fmt.Sprintf("%s%s%s", Italic, msg, Reset)
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

// Pad both sides of the string so the message is centered
func padCenter(s string, n int) string {
	if len(s) < n {
		s = fmt.Sprintf("%*s%s%*s", (n-len(s))/2, "", s, (n-len(s))/2, "")
	}
	return s
}

// Get the ECL default style log in string
func getDefaultStyleLog(msg message.LogMessage) string {
	pid := os.Getpid()

	// If the Name is empty, then set it to the default value
	if msg.Name == "" {
		msg.Name = "default"
	}

	return fmt.Sprintf(
		"%s %s %s %s %s - %s\n", // Format string
		colourize(msg.Color, "| "+bold(padMinWidthRight(msg.AppName, 12)+" |")), // Set name of app (min 12 characters)
		colourize(msg.Color, italic(padMinWidthRight(strconv.Itoa(pid), 6))),    // Add the process id
		colourize(White, msg.Time.Format(time.RFC3339)),                         // Add the time (time is white)
		colourize(msg.Color, bold(padMinWidthRight(msg.Level, 6))),              // Add the log level
		colourize(Yellow, padMinWidthRight("["+msg.Name+"]", 20)),               // Add the log name (name of the logger is yellow)
		colourize(msg.Color, msg.Msg),                                           // Add the message
	)
}

// Get the NestJS style log string
func getNestjsStyleLog(msg message.LogMessage) string {
	pid := os.Getpid()

	// If the Name is empty, then set it to the default value
	if msg.Name == "" {
		msg.Name = "default"
	}

	return fmt.Sprintf(
		"%s %-7s - %s %s %s %s\n",                                    // Format string
		colourize(msg.Color, "["+msg.AppName+"]"),                    // Set colour
		colourize(msg.Color, padMinWidthRight(strconv.Itoa(pid), 6)), // Add the process id
		colourize(White, msg.Time.Format("01/02/2006, 3:04:05 PM")),  // Add the time (time is white)
		colourize(msg.Color, padMinWidthLeft(msg.Level, 6)),          // Add the log level
		colourize(Yellow, "["+msg.Name+"]"),                          // Add the log name (name of the logger is yellow)
		colourize(msg.Color, msg.Msg),                                // Add the message
	)
}

// Print Spring style log
func getSpringStyleLog(msg message.LogMessage) string {
	pid := os.Getpid()
	thread := "main" // Thread is always main

	timeStr := msg.Time.Format(time.RFC3339)

	// Split the date-time
	date := timeStr[:10]
	time := timeStr[11:]

	return fmt.Sprintf(
		"%s %s %s --- %s %s %s\n",                           // <date-time>  <log level> <process id> --- [<thread>] <logger> : <message>
		colourize(White, date+" "+time),                     // Add the date-time (time is white)
		colourize(msg.Color, padMinWidthLeft(msg.Level, 6)), // Add the log level
		colourize(White, fmt.Sprintf("%d", pid)),            // Add the process id
		colourize(Yellow, "["+thread+"]"),                   // Add the thread
		colourize(Yellow, padMinWidthRight(msg.Name, 20)),   // Add the log name (name of the logger is yellow)
		colourize(msg.Color, msg.Msg),                       // Add the message
	)
}

func GetMessageOfStyle(msg message.LogMessage, logStyle LogStyle) string {
	// Get the log message in the given style
	switch logStyle {
	case NestJsStyle:
		return getNestjsStyleLog(msg)
	case SpringStyle:
		return getSpringStyleLog(msg)
	case DefaultStyle:
		return getDefaultStyleLog(msg)
	default:
		return getDefaultStyleLog(msg)
	}
}
