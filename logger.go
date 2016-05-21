package logplus

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Logger struct {
	timeFormat string
	logLevel   LogLevel
	mutex      *sync.Mutex
}

func NewLogger(tf string, l LogLevel) *Logger {
	return &Logger{
		timeFormat: tf,
		logLevel:   l,
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

func (logger *Logger) print(log ...interface{}) {
	logger.mutex.Lock()

	logger.printCallInfo(4)
	fmt.Print(log...)

	logger.mutex.Unlock()
}

func (logger *Logger) printf(format string, log ...interface{}) {
	logger.mutex.Lock()

	logger.printCallInfo(4)
	fmt.Printf(format, log...)

	logger.mutex.Unlock()
}

func (logger *Logger) println(log ...interface{}) {
	logger.mutex.Lock()

	logger.printCallInfo(4)
	fmt.Println(log...)

	logger.mutex.Unlock()
}

func (logger *Logger) coloredPrint(color Color, log ...interface{}) {
	logger.mutex.Lock()

	logger.printCallInfo(4)
	s := fmt.Sprint(log...)
	fmt.Print(concat("", color.String(), s, resetAll))

	logger.mutex.Unlock()
}

func (logger *Logger) coloredPrintf(color Color, format string, log ...interface{}) {
	logger.mutex.Lock()

	logger.printCallInfo(4)
	s := fmt.Sprintf(format, log...)
	fmt.Print(concat("", color.String(), s, resetAll))

	logger.mutex.Unlock()
}

func (logger *Logger) coloredPrintln(color Color, log ...interface{}) {
	logger.mutex.Lock()

	logger.printCallInfo(4)
	s := fmt.Sprint(log...)
	fmt.Println(concat("", color.String(), s, resetAll))

	logger.mutex.Unlock()
}

func (logger *Logger) leveledPrint(level LogLevel, log ...interface{}) {
	logger.mutex.Lock()

	if logger.isLogAvailable(level) {
		logger.printLogInfo(level)

		fmt.Print(log...)

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
		logger.printLogInfo(level)

		fmt.Printf(format, log...)

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
		logger.printLogInfo(level)

		fmt.Println(log...)

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

func (logger *Logger) printLogInfo(level LogLevel) {
	fmt.Print(concat("", level.Color().String(), level.String(), resetAll), " ")
	logger.printCallInfo(5)
}

func (logger *Logger) printCallInfo(depth int) {
	t := time.Now().Format(logger.timeFormat)
	info := getCallInfo(depth)
	fmt.Printf("%s [%s] <%d> ", t, info, os.Getpid())
}

func (logger *Logger) isLogAvailable(level LogLevel) bool {
	if level <= logger.logLevel {
		return true
	}
	return false
}

func concat(sep string, strs ...string) string {
	var result = make([]byte, 0, 100)
	for i, _ := range strs {
		result = append(result, strs[i]...)
		if i < len(strs)-1 {
			result = append(result, sep...)
		}
	}
	return string(result)
}
