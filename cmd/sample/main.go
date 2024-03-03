// This is a sample entrypoint to test the package
package main

import (
	logger "github.com/jhseong7/ecl"
)

func main() {
	l := logger.NewLogger(logger.LoggerOption{
		Name: "test",
	})

	l.Log("Hello, World!")
}