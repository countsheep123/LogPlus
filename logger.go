package logplus

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	output     io.Writer
	timeFormat string
	logLevel   LogLevel
	formatter  Formatter
	mutex      *sync.Mutex
}

func NewLogger(o io.Writer, tf string, l LogLevel, f Formatter) *Logger {
	return &Logger{
		output:     o,
		timeFormat: tf,
		logLevel:   l,
		formatter:  f,
		mutex:      new(sync.Mutex),
	}
}

func (logger *Logger) setTimeFormat(tf string) {
	logger.mutex.Lock()
	logger.timeFormat = tf
	logger.mutex.Unlock()
}

func (logger *Logger) setLevel(level LogLevel) {
	logger.mutex.Lock()
	logger.logLevel = level
	logger.mutex.Unlock()
}

func (logger *Logger) setOutput(out io.Writer) {
	logger.mutex.Lock()
	logger.output = out
	logger.mutex.Unlock()
}

func (logger *Logger) setFormatter(formatter Formatter) {
	logger.mutex.Lock()
	logger.formatter = formatter
	logger.mutex.Unlock()
}

func (logger *Logger) print(log ...interface{}) {
	logger.mutex.Lock()

	t := time.Now().Format(logger.timeFormat)
	caller := getCallInfo(3)
	pid := os.Getpid()

	msg := NewLogMessage(NoColor, log...)
	entry := NewLogEntry(none, t, caller, pid, msg)
	fmtLog := logger.formatter.Format(entry)

	//info := fmt.Sprintf("%s [%s] <%d>", t, caller, pid)
	//s := fmt.Sprint(log...)
	//fmt.Fprint(logger.output, concat(": ", info, s))
	fmt.Fprint(logger.output, fmtLog)

	logger.mutex.Unlock()
}

func (logger *Logger) printf(format string, log ...interface{}) {
	logger.mutex.Lock()

	t := time.Now().Format(logger.timeFormat)
	caller := getCallInfo(3)
	pid := os.Getpid()

	msg := NewFormattedLogMessage(NoColor, format, log...)
	entry := NewLogEntry(none, t, caller, pid, msg)
	fmtLog := logger.formatter.Format(entry)

	//info := fmt.Sprintf("%s [%s] <%d>", t, caller, pid)
	//s := fmt.Sprintf(format, log...)
	//fmt.Fprintf(logger.output, concat(": ", info, s))
	fmt.Fprint(logger.output, fmtLog)

	logger.mutex.Unlock()
}

func (logger *Logger) println(log ...interface{}) {
	logger.mutex.Lock()

	t := time.Now().Format(logger.timeFormat)
	caller := getCallInfo(3)
	pid := os.Getpid()

	msg := NewLogMessage(NoColor, log...)
	entry := NewLogEntry(none, t, caller, pid, msg)
	fmtLog := logger.formatter.Format(entry)

	//info := fmt.Sprintf("%s [%s] <%d>", t, caller, pid)
	//s := fmt.Sprint(log...)
	//fmt.Fprintln(logger.output, concat(": ", info, s))
	fmt.Fprintln(logger.output, fmtLog)

	logger.mutex.Unlock()
}

func (logger *Logger) coloredPrint(color Color, log ...interface{}) {
	logger.mutex.Lock()

	t := time.Now().Format(logger.timeFormat)
	caller := getCallInfo(3)
	pid := os.Getpid()

	msg := NewLogMessage(color, log...)
	entry := NewLogEntry(none, t, caller, pid, msg)
	fmtLog := logger.formatter.Format(entry)

	//info := fmt.Sprintf("%s [%s] <%d>", t, caller, pid)
	//s := fmt.Sprint(log...)
	//fmt.Fprint(logger.output, concat(": ", info, concat("", color.String(), s, resetAll)))
	fmt.Fprint(logger.output, fmtLog)

	logger.mutex.Unlock()
}

func (logger *Logger) coloredPrintf(color Color, format string, log ...interface{}) {
	logger.mutex.Lock()

	t := time.Now().Format(logger.timeFormat)
	caller := getCallInfo(3)
	pid := os.Getpid()

	msg := NewFormattedLogMessage(color, format, log...)
	entry := NewLogEntry(none, t, caller, pid, msg)
	fmtLog := logger.formatter.Format(entry)

	//info := fmt.Sprintf("%s [%s] <%d>", t, caller, pid)
	//s := fmt.Sprintf(format, log...)
	//fmt.Fprint(logger.output, concat(": ", info, concat("", color.String(), s, resetAll)))

	fmt.Fprint(logger.output, fmtLog)

	logger.mutex.Unlock()
}

func (logger *Logger) coloredPrintln(color Color, log ...interface{}) {
	logger.mutex.Lock()

	t := time.Now().Format(logger.timeFormat)
	caller := getCallInfo(3)
	pid := os.Getpid()

	msg := NewLogMessage(color, log...)
	entry := NewLogEntry(none, t, caller, pid, msg)
	fmtLog := logger.formatter.Format(entry)

	//info := fmt.Sprintf("%s [%s] <%d>", t, caller, pid)
	//s := fmt.Sprint(log...)
	//fmt.Fprintln(logger.output, concat(": ", info, concat("", color.String(), s, resetAll)))

	fmt.Fprintln(logger.output, fmtLog)

	logger.mutex.Unlock()
}

func (logger *Logger) leveledPrint(level LogLevel, log ...interface{}) {
	logger.mutex.Lock()

	if logger.isLogAvailable(level) {
		t := time.Now().Format(logger.timeFormat)
		caller := getCallInfo(4)
		pid := os.Getpid()

		msg := NewLogMessage(NoColor, log...)
		entry := NewLogEntry(level, t, caller, pid, msg)
		fmtLog := logger.formatter.Format(entry)

		fmt.Fprint(logger.output, fmtLog)

		switch level {
		case FATAL:
			os.Exit(1)
		case PANIC:
			panic(fmt.Sprint(log...))
		default:
		}
	}

	logger.mutex.Unlock()
}

func (logger *Logger) leveledPrintf(level LogLevel, format string, log ...interface{}) {
	logger.mutex.Lock()

	if logger.isLogAvailable(level) {
		t := time.Now().Format(logger.timeFormat)
		caller := getCallInfo(4)
		pid := os.Getpid()

		msg := NewFormattedLogMessage(NoColor, format, log...)
		entry := NewLogEntry(level, t, caller, pid, msg)
		fmtLog := logger.formatter.Format(entry)

		fmt.Fprint(logger.output, fmtLog)

		switch level {
		case FATAL:
			os.Exit(1)
		case PANIC:
			panic(fmt.Sprintf(format, log...))
		default:
		}
	}

	logger.mutex.Unlock()
}

func (logger *Logger) leveledPrintln(level LogLevel, log ...interface{}) {
	logger.mutex.Lock()

	if logger.isLogAvailable(level) {
		t := time.Now().Format(logger.timeFormat)
		caller := getCallInfo(4)
		pid := os.Getpid()

		msg := NewLogMessage(NoColor, log...)
		entry := NewLogEntry(level, t, caller, pid, msg)
		fmtLog := logger.formatter.Format(entry)

		fmt.Fprintln(logger.output, fmtLog)

		switch level {
		case FATAL:
			os.Exit(1)
		case PANIC:
			panic(fmt.Sprintln(log...))
		default:
		}
	}

	logger.mutex.Unlock()
}

func (logger *Logger) isLogAvailable(level LogLevel) bool {
	if level <= logger.logLevel {
		return true
	}
	return false
}
