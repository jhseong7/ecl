package stream

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"sync"
	"time"

	"github.com/jhseong7/ecl/message"
	"github.com/jhseong7/ecl/style"
)

type (
	FileLogStreamOption struct {
		LogDirectory  string
		FileName      string
		FileRollover  bool
		MaxFileSizeKb int
		LogStyle      style.LogStyle
	}

	FileLogStream struct {
		ILogStream

		// Copy of the initial options
		options FileLogStreamOption

		// File pointer
		file *os.File

		// Mutex to prevent multiple writes at the same time
		mutex *sync.Mutex
	}
)

var (
	terminalStyleRegex = regexp.MustCompile("\x1b\\[[0-9;]*m")
)

func (s *FileLogStream) getLogFileName(prefix, date string) string {
	if s.options.FileRollover {
		return fmt.Sprintf("%s.%s.log", prefix, date)
	}

	return fmt.Sprintf("%s.log", prefix)
}

// NOTE: Use this later when rollover is implemented
// func (s *FileLogStream) getFileSizeKb() int64 {
// 	fileInfo, err := s.file.Stat()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return fileInfo.Size() / 1024
// }

func (s *FileLogStream) createNewLog() {
	// Close the file
	s.file.Close()

	// If the LogDirectory does not exist, create it (recursively)
	if _, err := os.Stat(s.options.LogDirectory); os.IsNotExist(err) {
		os.MkdirAll(s.options.LogDirectory, 0755)
	}

	// Open a new file
	f, err := os.OpenFile(
		path.Join(s.options.LogDirectory, s.getLogFileName(s.options.FileName, time.Now().Format("2006-01-02"))),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		panic(err)
	}

	s.file = f
}

func (s *FileLogStream) getFilePointer() *os.File {
	// Get the current time, and if the current file is not the same date, close the file and open a new one with the current date
	currentDate := time.Now().Format("2006-01-02")
	logFileName := s.getLogFileName(s.options.FileName, currentDate)

	// If there is no file pointer, open a new file
	// If the file is not the same date, close the file and open a new one
	if s.file == nil || s.file.Name() != logFileName {
		s.createNewLog()
	}

	return s.file
}

func (s *FileLogStream) Write(msg message.LogMessage) {
	// Use a mutex to prevent multiple writes at the same time
	s.mutex.Lock()
	defer s.mutex.Unlock()

	file := s.getFilePointer()

	// Get the message --> and remove all terminal styleing
	msgStr := style.GetMessageOfStyle(msg, s.options.LogStyle)
	msgStr = terminalStyleRegex.ReplaceAllString(msgStr, "")

	// Append the message to the file (with a new line)
	fmt.Fprint(file, msgStr)
}

func NewFileLogStream(option FileLogStreamOption) *FileLogStream {
	// Check if all options are given
	if option.LogDirectory == "" || option.FileName == "" {
		panic("NewFileLogStream: LogDirectory and FileName must be given")
	}

	return &FileLogStream{
		options: option,
		mutex:   &sync.Mutex{},
	}
}
