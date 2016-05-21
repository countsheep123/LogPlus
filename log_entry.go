package logplus

type LogEntry struct {
	Level   LogLevel
	Time    string
	Caller  *CallInfo
	Pid     int
	Message *LogMessage
}

func NewLogEntry(level LogLevel, t string, caller *CallInfo, pid int, msg *LogMessage) *LogEntry {
	return &LogEntry{
		Level:   level,
		Time:    t,
		Caller:  caller,
		Pid:     pid,
		Message: msg,
	}
}
