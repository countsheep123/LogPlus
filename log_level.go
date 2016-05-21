package logplus

type LogLevel int

const (
	PANIC LogLevel = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	none
)

func (level LogLevel) String() string {
	switch level {
	case PANIC:
		return "PANIC"
	case FATAL:
		return "FATAL"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	}
	panic("Unknown value")
}

func (level LogLevel) Color() Color {
	switch level {
	case PANIC:
		return BackgroundRed
	case FATAL:
		return BackgroundRed
	case ERROR:
		return BackgroundMagenta
	case WARN:
		return BackgroundYellow
	case INFO:
		return BackgroundGreen
	case DEBUG:
		return BackgroundCyan
	}
	panic("Unknown value")
}
