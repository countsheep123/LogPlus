package logplus

import (
	"encoding/json"
)

type JsonFormatter struct {
}

func (jf *JsonFormatter) Format(entry *LogEntry) string {
	switch entry.Level {
	case none:
		jsonMap := map[string]interface{}{
			"time":   entry.Time,
			"caller": entry.Caller.String(),
			"pid":    entry.Pid,
			"log":    entry.Message.message,
		}
		jsonBytes, err := json.Marshal(jsonMap)
		if err != nil {
			return ""
		}
		return string(jsonBytes)
	default:
		jsonMap := map[string]interface{}{
			"level":  entry.Level,
			"time":   entry.Time,
			"caller": entry.Caller.String(),
			"pid":    entry.Pid,
			"log":    entry.Message.message,
		}
		jsonBytes, err := json.Marshal(jsonMap)
		if err != nil {
			return ""
		}
		return string(jsonBytes)
	}
}
