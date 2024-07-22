package message

import "time"

type (
	LogMessage struct {
		AppName string
		Time    time.Time
		Name    string
		Color   string
		Level   string
		Msg     string
	}
)
