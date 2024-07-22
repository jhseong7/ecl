package ecl

import (
	"github.com/jhseong7/ecl/logger"
	"github.com/jhseong7/ecl/stream"
	"github.com/jhseong7/ecl/style"
)

type (
	ILogStream = stream.ILogStream

	LoggerOption = logger.LoggerOption
	Logger       = logger.Logger

	LogLevel = logger.LogLevel

	LogStyle = style.LogStyle
)

const (
	NestJsStyle  = style.NestJsStyle
	SpringStyle  = style.SpringStyle
	DefaultStyle = style.DefaultStyle

	All   = logger.All
	Trace = logger.Trace
	Debug = logger.Debug
	Info  = logger.Info
	Warn  = logger.Warn
	Error = logger.Error
)

func NewLogger(o LoggerOption) Logger {
	return logger.NewLogger(o)
}

func SetAppName(name string) {
	logger.SetAppName(name)
}

func AddGlobalExtraStream(streams []ILogStream) {
	logger.AddGlobalExtraStream(streams)
}

func SetLogStyle(style LogStyle) {
	logger.SetLogStyle(style)
}

// Set the log level for the app. This will be used for all loggers.
func SetLogLevel(level LogLevel) {
	logger.SetLogLevel(level)
}
