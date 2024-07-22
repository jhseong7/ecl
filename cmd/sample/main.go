// This is a sample entrypoint to test the package
package main

import (
	logger "github.com/jhseong7/ecl"
)

func main() {
	l := logger.NewLogger(logger.LoggerOption{
		Name: "test",
	})

	l2 := logger.NewLogger(logger.LoggerOption{
		Name: "OtherService",
	})

	l3 := logger.NewLogger(logger.LoggerOption{})

	l.Log("Hello, World!")
	l2.Log("Hello, OtherService!")
	l.Warn("This is a warning")
	l2.Error("This is an error")
	l.Trace("This is a trace")
	l2.Debug("This is a debug")
	l3.Log("Log with no name")
}
