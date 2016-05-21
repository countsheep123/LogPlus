package logplus

import (
	"fmt"
)

type TextFormatter struct {
}

func (tf *TextFormatter) Format(entry *LogEntry) string {
	switch entry.Level {
	case none:
		if entry.Message.color == NoColor {
			return fmt.Sprintf("%s [%s] <%d>: %s", entry.Time, entry.Caller, entry.Pid, entry.Message.message)
		} else {
			return fmt.Sprintf("%s [%s] <%d>: %v%s%s", entry.Time, entry.Caller, entry.Pid, entry.Message.color, entry.Message.message, resetAll)
		}
	default:
		if entry.Message.color == NoColor {
			return fmt.Sprintf("%v%-5s%s %s [%s] <%d>: %s", entry.Level.Color(), entry.Level.String(), resetAll, entry.Time, entry.Caller, entry.Pid, entry.Message.message)
		} else {
			return fmt.Sprintf("%v%-5s%s %s [%s] <%d>: %v%s%s", entry.Level.Color(), entry.Level.String(), resetAll, entry.Time, entry.Caller, entry.Pid, entry.Message.color, entry.Message.message, resetAll)
		}
	}
}
