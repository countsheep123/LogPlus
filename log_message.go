package logplus

import (
	"fmt"
)

type LogMessage struct {
	message string
	color   Color
}

func NewFormattedLogMessage(color Color, format string, log ...interface{}) *LogMessage {
	return &LogMessage{
		message: fmt.Sprintf(format, log...),
		color:   color,
	}
}

func NewLogMessage(color Color, log ...interface{}) *LogMessage {
	return &LogMessage{
		message: fmt.Sprint(log...),
		color:   color,
	}
}
