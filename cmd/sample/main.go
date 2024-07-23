// This is a sample entrypoint to test the package
package main

import (
	"github.com/jhseong7/ecl"
)

func main() {
	ecl.SetLogLevel(ecl.All)
	ecl.SetAppName("ExampleApp")

	l := ecl.NewLogger(ecl.LoggerOption{
		Name: "Example",
	})

	l2 := ecl.NewLogger(ecl.LoggerOption{
		Name: "OtherService",
	})

	l3 := ecl.NewLogger(ecl.LoggerOption{
		AppName: "Another",
		Name:    "Service",
	})

	l.Log("Hello, World!")
	l2.Log("Hello, OtherService!")
	l.Warn("This is a warning")
	l.Error("This is an error")
	l.Trace("This is a trace")
	l.Debug("This is a debug")
	l.Info("This is a info")
	l3.Log("Logger with alternate app name")
}
