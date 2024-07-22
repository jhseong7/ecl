package stream

import (
	"fmt"
	"os"

	"github.com/jhseong7/ecl/message"
	"github.com/jhseong7/ecl/style"
)

type (
	StdOutStream struct {
		ILogStream
		logStyle style.LogStyle
	}

	StdOutStreamOption struct {
		LogStyle style.LogStyle
	}
)

func (s *StdOutStream) Write(msg message.LogMessage) {
	fmt.Print(style.GetMessageOfStyle(msg, s.logStyle))
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
	switch logStyle {
	case string(style.NestJsStyle):
		return &StdOutStream{
			logStyle: style.NestJsStyle,
		}
	case string(style.SpringStyle):
		return &StdOutStream{
			logStyle: style.SpringStyle,
		}
	}

	return &StdOutStream{}
}
