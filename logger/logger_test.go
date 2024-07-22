package logger_test

import (
	"testing"

	"github.com/jhseong7/ecl/logger"
	"github.com/jhseong7/ecl/message"
	"github.com/jhseong7/ecl/stream"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestStream struct {
	stream.ILogStream
	LastMessage message.LogMessage
}

func (s *TestStream) Write(msg message.LogMessage) {
	s.LastMessage = msg
}

var _ = Describe("Coloured Logger", func() {
	// Setup the logger
	ts := &TestStream{}
	l := logger.NewLogger(logger.LoggerOption{
		Name:   "test",
		Silent: true,
		ExtraStreams: []stream.ILogStream{
			ts,
		},
	})

	// Test the logger
	It("Test Log", func() {
		const msg = "Hello, world!"
		l.Log(msg)

		// Check the contents
		m := ts.LastMessage

		// Check the message
		Expect(m.Msg).To(Equal(msg))

		// Check the level
		Expect(m.Level).To(Equal("LOG"))
	})

	It("Test Logf", func() {
		const msg = "Hello, world!"
		l.Logf("%s", msg)

		// Check the contents
		m := ts.LastMessage

		// Check the message
		Expect(m.Msg).To(Equal(msg))

		// Check the level
		Expect(m.Level).To(Equal("LOG"))
	})

	It("Test Trace", func() {
		const msg = "Hello, world!"
		l.Trace(msg)

		// Check the contents
		m := ts.LastMessage

		// Check the message
		Expect(m.Msg).To(Equal(msg))

		// Check the level
		Expect(m.Level).To(Equal("TRACE"))
	})

	It("Test Tracef", func() {
		const msg = "Hello, world!"
		l.Tracef("%s", msg)

		// Check the contents
		m := ts.LastMessage

		// Check the message
		Expect(m.Msg).To(Equal(msg))

		// Check the level
		Expect(m.Level).To(Equal("TRACE"))
	})

	It("Test WARN", func() {
		const msg = "Hello, world!"
		l.Warn(msg)

		// Check the contents
		m := ts.LastMessage

		// Check the message
		Expect(m.Msg).To(Equal(msg))

		// Check the level
		Expect(m.Level).To(Equal("WARN"))
	})

	It("Test Error", func() {
		const msg = "Hello, world!"
		l.Error(msg)

		// Check the contents
		m := ts.LastMessage

		// Check the message
		Expect(m.Msg).To(Equal(msg))

		// Check the level
		Expect(m.Level).To(Equal("ERROR"))
	})
})

func TestLogger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Logger Suite")
}
