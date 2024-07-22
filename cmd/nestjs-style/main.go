// This is a sample entrypoint to test the package
package main

import (
	"github.com/jhseong7/ecl"
	"github.com/jhseong7/ecl/style"
)

func main() {
	ecl.SetLogStyle(style.NestJsStyle)

	l := ecl.NewLogger(ecl.LoggerOption{
		Name: "test",
	})

	l2 := ecl.NewLogger(ecl.LoggerOption{
		Name: "OtherService",
	})

	l.Log("Hello, World!")
	l2.Log("Hello, OtherService!")
	l.Warn("This is a warning")
	l2.Error("This is an error")
	l.Trace("This is a trace")
	l2.Debug("This is a debug")
}
