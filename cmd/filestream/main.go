// This is a sample entrypoint to test the package
package main

import (
	"github.com/jhseong7/ecl"
	"github.com/jhseong7/ecl/stream"
)

func main() {
	ecl.SetLogLevel(ecl.All)

	ecl.AddGlobalExtraStream([]ecl.ILogStream{
		stream.NewFileLogStream(stream.FileLogStreamOption{
			LogDirectory: "temp",
			FileName:     "default-style",
		}),
		stream.NewFileLogStream(stream.FileLogStreamOption{
			LogDirectory: "temp",
			FileName:     "spring-style",
			LogStyle:     ecl.SpringStyle,
		}),
		stream.NewFileLogStream(stream.FileLogStreamOption{
			LogDirectory: "temp",
			FileName:     "nestjs-style",
			LogStyle:     ecl.NestJsStyle,
		}),
	})

	l := ecl.NewLogger(ecl.LoggerOption{
		Name: "DefaultService",
	})

	l2 := ecl.NewLogger(ecl.LoggerOption{
		Name:     "OtherService",
		LogStyle: ecl.SpringStyle,
	})

	l3 := ecl.NewLogger(ecl.LoggerOption{
		Name:     "OtherService2",
		LogStyle: ecl.NestJsStyle,
	})

	l.Log("Hello, World!")
	l2.Log("Hello, OtherService!")

	l.Warn("This is a warning")
	l.Error("This is an error")
	l.Trace("This is a trace")
	l.Debug("This is a debug")

	l2.Warn("This is a warning")
	l2.Error("This is an error")
	l2.Trace("This is a trace")
	l2.Debug("This is a debug")

	l3.Warn("This is a warning")
	l3.Error("This is an error")
	l3.Trace("This is a trace")
	l3.Debug("This is a debug")
}
