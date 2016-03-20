//	logplus.Logln("this is example")

package logplus

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	resetAll = "\033[0m"
)

var (
	timeFormat = time.RFC3339
	logLevel   = INFO
)

type Color int

const (
	ForegroundDefault Color = iota
	ForegroundBlack
	ForegroundRed
	ForegroundGreen
	ForegroundYellow
	ForegroundBlue
	ForegroundMagenta
	ForegroundCyan
	ForegroundWhite

	BackgroundDefault
	BackgroundBlack
	BackgroundRed
	BackgroundGreen
	BackgroundYellow
	BackgroundBlue
	BackgroundMagenta
	BackgroundCyan
	BackgroundWhite
)

func (color Color) String() string {
	switch color {
	case ForegroundDefault:
		return "\033[39m"
	case ForegroundBlack:
		return "\033[30m"
	case ForegroundRed:
		return "\033[31m"
	case ForegroundGreen:
		return "\033[32m"
	case ForegroundYellow:
		return "\033[33m"
	case ForegroundBlue:
		return "\033[34m"
	case ForegroundMagenta:
		return "\033[35m"
	case ForegroundCyan:
		return "\033[36m"
	case ForegroundWhite:
		return "\033[97m"
	case BackgroundDefault:
		return "\033[49m"
	case BackgroundBlack:
		return "\033[40m"
	case BackgroundRed:
		return "\033[41m"
	case BackgroundGreen:
		return "\033[42m"
	case BackgroundYellow:
		return "\033[43m"
	case BackgroundBlue:
		return "\033[44m"
	case BackgroundMagenta:
		return "\033[45m"
	case BackgroundCyan:
		return "\033[46m"
	case BackgroundWhite:
		return "\033[107m"
	}
	panic("Unknown value")
}

type LogLevel int

const (
	FATAL LogLevel = iota
	ERROR
	WARN
	INFO
	DEBUG
)

func (level LogLevel) String() string {
	switch level {
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

type CallInfo struct {
	PackageName string
	FileName    string
	FuncName    string
	Line        int
}

func (ci CallInfo) String() string {
	return fmt.Sprintf("%s#%s (%s:%d)", ci.PackageName, ci.FuncName, ci.FileName, ci.Line)
}

func SetTimeFormat(tf string) {
	timeFormat = tf
}

func SetLevel(level LogLevel) {
	logLevel = level
}

func isLogAvailable(level LogLevel) bool {
	if logLevel >= level {
		return true
	}
	return false
}

func Log(log ...interface{}) {
	info := getCallInfo()
	fmt.Printf(" %s [%s] <%d> ", time.Now().Format(timeFormat), info, os.Getpid)
	fmt.Print(log...)
}

func Logf(format string, log ...interface{}) {
	info := getCallInfo()
	fmt.Printf(" %s [%s] <%d> ", time.Now().Format(timeFormat), info, os.Getpid)
	fmt.Printf(format, log...)
}

func Logln(log ...interface{}) {
	info := getCallInfo()
	fmt.Printf(" %s [%s] <%d> ", time.Now().Format(timeFormat), info, os.Getpid)
	fmt.Println(log...)
}

func ColoredLog(color Color, log ...interface{}) {
	fmt.Print(color)
	Log(log...)
	fmt.Print(resetAll)
}

func ColoredLogf(color Color, format string, log ...interface{}) {
	fmt.Print(color)
	Logf(format, log...)
	fmt.Print(resetAll)
}

func ColoredLogln(color Color, log ...interface{}) {
	fmt.Print(color)
	Logln(log...)
	fmt.Print(resetAll)
}

func FatalLog(log ...interface{}) {
	if isLogAvailable(FATAL) {
		fmt.Print(BackgroundRed, FATAL, resetAll)
		Log(log...)
	}
}

func FatalLogf(format string, log ...interface{}) {
	if isLogAvailable(FATAL) {
		fmt.Print(BackgroundRed, FATAL, resetAll)
		Logf(format, log...)
	}
}

func FatalLogln(log ...interface{}) {
	if isLogAvailable(FATAL) {
		fmt.Print(BackgroundRed, FATAL, resetAll)
		Logln(log...)
	}
}

func ErrorLog(log ...interface{}) {
	if isLogAvailable(ERROR) {
		fmt.Print(BackgroundMagenta, ERROR, resetAll)
		Log(log...)
	}
}

func ErrorLogf(format string, log ...interface{}) {
	if isLogAvailable(ERROR) {
		fmt.Print(BackgroundMagenta, ERROR, resetAll)
		Logf(format, log...)
	}
}

func ErrorLogln(log ...interface{}) {
	if isLogAvailable(ERROR) {
		fmt.Print(BackgroundMagenta, ERROR, resetAll)
		Logln(log...)
	}
}

func WarnLog(log ...interface{}) {
	if isLogAvailable(WARN) {
		fmt.Print(BackgroundYellow, WARN, resetAll, " ")
		Log(log...)
	}
}

func WarnLogf(format string, log ...interface{}) {
	if isLogAvailable(WARN) {
		fmt.Print(BackgroundYellow, WARN, resetAll, " ")
		Logf(format, log...)
	}
}

func WarnLogln(log ...interface{}) {
	if isLogAvailable(WARN) {
		fmt.Print(BackgroundYellow, WARN, resetAll, " ")
		Logln(log...)
	}
}

func InfoLog(log ...interface{}) {
	if isLogAvailable(INFO) {
		fmt.Print(BackgroundGreen, INFO, resetAll, " ")
		Log(log...)
	}
}

func InfoLogf(format string, log ...interface{}) {
	if isLogAvailable(INFO) {
		fmt.Print(BackgroundGreen, INFO, resetAll, " ")
		Logf(format, log...)
	}
}

func InfoLogln(log ...interface{}) {
	if isLogAvailable(INFO) {
		fmt.Print(BackgroundGreen, INFO, resetAll, " ")
		Logln(log...)
	}
}

func DebugLog(log ...interface{}) {
	if isLogAvailable(DEBUG) {
		fmt.Print(BackgroundCyan, DEBUG, resetAll)
		Log(log...)
	}
}

func DebugLogf(format string, log ...interface{}) {
	if isLogAvailable(DEBUG) {
		fmt.Print(BackgroundCyan, DEBUG, resetAll)
		Logf(format, log...)
	}
}

func DebugLogln(log ...interface{}) {
	if isLogAvailable(DEBUG) {
		fmt.Print(BackgroundCyan, DEBUG, resetAll)
		Logln(log...)
	}
}

func getCallInfo() *CallInfo {
	pc, filePath, line, _ := runtime.Caller(2)
	_, fileName := path.Split(filePath)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]
	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &CallInfo{
		PackageName: packageName,
		FileName:    fileName,
		FuncName:    funcName,
		Line:        line,
	}
}
