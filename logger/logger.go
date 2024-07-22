package logger

import (
	"fmt"
	"log"
	"time"

	"github.com/jhseong7/ecl/message"
	"github.com/jhseong7/ecl/stream"
	"github.com/jhseong7/ecl/style"
)

type (
	LoggerOption struct {
		// Name of the logger. This will be used as the logger's name to identify the logger
		Name string

		// If true, the logger will not print anything to stdout
		Silent bool

		// Extra streams to write to
		ExtraStreams []stream.ILogStream

		// Log style. Default is NestJsStyle
		LogStyle style.LogStyle

		// Local log level. Default is All
		LogLevel LogLevel

		// Local App name. If set, this name will be added to all log messages as a prefix.
		AppName string
	}

	Logger interface {
		Log(msg string)
		Trace(msg string)
		Debug(msg string)
		Info(msg string)
		Warn(msg string)
		Error(msg string)
		Fatal(msg string)
		Panic(msg string)

		// format functions
		Logf(format string, args ...interface{})
		Tracef(format string, args ...interface{})
		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalf(format string, args ...interface{})
		Panicf(format string, args ...interface{})
	}

	LoggerImpl struct {
		Logger
		Streams  []stream.ILogStream
		name     string
		loglevel LogLevel
		appName  string
	}

	LogLevel int
)

const (
	All LogLevel = iota
	Trace
	Debug
	Info
	Warn
	Error

	// No options for Fatal and Panic (always print)
)

var (
	// Global prefix for the logger. Default is "GoApp"
	globalPrefix = "GoApp"

	// Global extra streams
	globalExtraStreams []stream.ILogStream

	// global logStyle
	globalLogStyle style.LogStyle = style.DefaultStyle

	globalLoglevel LogLevel
)

func NewLogger(o LoggerOption) Logger {
	var ss []stream.ILogStream

	if !o.Silent {
		var logStyle style.LogStyle
		if o.LogStyle != "" {
			logStyle = o.LogStyle
		} else {
			logStyle = globalLogStyle
		}

		ss = append(ss, stream.NewStdOutStream(stream.StdOutStreamOption{
			LogStyle: logStyle,
		}))
	}

	// Append global extra streams
	ss = append(ss, globalExtraStreams...)

	// Append extra streams
	ss = append(ss, o.ExtraStreams...)

	// Set loglevel
	var loglevel LogLevel
	if o.LogLevel != 0 {
		loglevel = o.LogLevel
	} else {
		loglevel = globalLoglevel
	}

	// Set appName
	var appName string
	if o.AppName != "" {
		appName = o.AppName
	} else {
		appName = globalPrefix
	}

	return &LoggerImpl{
		name:     o.Name,
		Streams:  ss,
		loglevel: loglevel,
		appName:  appName,
	}
}

// Set the log level for the app. This will be used for all loggers.
func SetLogLevel(level LogLevel) {
	globalLoglevel = level
}

// Set the global prefix for the logger. If set, this name will be added to all log messages as a prefix.
func SetAppName(prefix string) {
	globalPrefix = prefix
}

// Set the global log style for the logger. If set, this style will be used for all log messages.
func SetLogStyle(style style.LogStyle) {
	globalLogStyle = style
}

// Add the streams to the global extra streams. This will be added to all loggers.
func AddGlobalExtraStream(streams []stream.ILogStream) {
	globalExtraStreams = append(globalExtraStreams, streams...)
}

func (l *LoggerImpl) writeToStream(color, logLevel, msg string) {
	// Get the current time here so that all streams have the same time
	ct := time.Now()

	// For all streams
	for _, stream := range l.Streams {
		// Write the log message
		stream.Write(message.LogMessage{
			AppName: l.appName,
			Name:    l.name,
			Time:    ct,
			Color:   color,
			Level:   logLevel,
			Msg:     msg,
		})
	}

}

func (l *LoggerImpl) logWithColorf(color, logLevel, format string, args ...interface{}) {
	l.writeToStream(color, logLevel, fmt.Sprintf(format, args...))
}

func (l *LoggerImpl) Log(msg string) {
	// Green
	l.writeToStream(style.Green, "LOG", msg)
}

// Formatted Log log. Use this like fmt.Printf
func (l *LoggerImpl) Logf(format string, args ...interface{}) {
	// Green
	l.logWithColorf(style.Green, "LOG", format, args...)
}

func (l *LoggerImpl) Trace(msg string) {
	if l.loglevel > Trace {
		return
	}

	l.writeToStream(style.Purple, "TRACE", msg)
}

// Formatted Trace log. Use this like fmt.Printf
func (l *LoggerImpl) Tracef(format string, args ...interface{}) {
	if l.loglevel > Trace {
		return
	}

	l.logWithColorf(style.Purple, "TRACE", format, args...)
}

func (l *LoggerImpl) Debug(msg string) {
	if l.loglevel > Debug {
		return
	}

	// Blue
	l.writeToStream(style.Blue, "DEBUG", msg)
}

// Formatted Debug log. Use this like fmt.Printf
func (l *LoggerImpl) Debugf(format string, args ...interface{}) {
	if l.loglevel > Debug {
		return
	}

	// Blue
	l.logWithColorf(style.Blue, "DEBUG", format, args...)
}

func (l *LoggerImpl) Info(msg string) {
	if l.loglevel > Info {
		return
	}

	// Green
	l.writeToStream(style.Cyan, "INFO", msg)
}

// Formatted Log log. Use this like fmt.Printf
func (l *LoggerImpl) Infof(format string, args ...interface{}) {
	if l.loglevel > Info {
		return
	}

	// Green
	l.logWithColorf(style.Cyan, "INFO", format, args...)
}

// Warn Log
func (l *LoggerImpl) Warn(msg string) {
	if l.loglevel > Warn {
		return
	}

	// Yellow
	l.writeToStream(style.Yellow, "WARN", msg)
}

// Formatted Warn log. Use this like fmt.Printf
func (l *LoggerImpl) Warnf(format string, args ...interface{}) {
	if l.loglevel > Warn {
		return
	}

	// Yellow
	l.logWithColorf(style.Yellow, "WARN", format, args...)
}

func (l *LoggerImpl) Error(msg string) {
	if l.loglevel > Error {
		return
	}

	// Red
	l.writeToStream(style.Red, "ERROR", msg)
}

// Formatted Error log. Use this like fmt.Printf
func (l *LoggerImpl) Errorf(format string, args ...interface{}) {
	if l.loglevel > Error {
		return
	}

	// Red
	l.logWithColorf(style.Red, "ERROR", format, args...)
}

func (l *LoggerImpl) Fatal(msg string) {
	// Red + Fatal + Exit(1)
	l.writeToStream(style.Red, "FATAL", msg)
	log.Fatal(msg)
}

// Formatted Fatal log. Use this like fmt.Printf
func (l *LoggerImpl) Fatalf(format string, args ...interface{}) {
	// Red + Fatal + Exit(1)
	l.logWithColorf(style.Red, "FATAL", format, args...)
	log.Fatalf(format, args...)
}

func (l *LoggerImpl) Panic(msg string) {
	// Red + Panic + Exit(1)
	l.writeToStream(style.Red, "PANIC", msg)
	panic(msg)
}

// Formatted Panic log. Use this like fmt.Printf
func (l *LoggerImpl) Panicf(format string, args ...interface{}) {
	// Red + Panic + Exit(1)
	l.logWithColorf(style.Red, "PANIC", format, args...)
	panic(fmt.Sprintf(format, args...))
}
