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
	if level >= logLevel {
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

func Colored(color Color, log ...interface{}) {
	fmt.Print(color)
	Log(log...)
	fmt.Print(resetAll)
}

func Coloredf(color Color, format string, log ...interface{}) {
	fmt.Print(color)
	Logf(format, log...)
	fmt.Print(resetAll)
}

func Coloredln(color Color, log ...interface{}) {
	fmt.Print(color)
	Logln(log...)
	fmt.Print(resetAll)
}

func Fatal(log ...interface{}) {
	if isLogAvailable(FATAL) {
		fmt.Print(concat("", BackgroundRed.String(), FATAL.String(), resetAll))
		Log(log...)
	}
}

func Fatalf(format string, log ...interface{}) {
	if isLogAvailable(FATAL) {
		fmt.Print(concat("", BackgroundRed.String(), FATAL.String(), resetAll))
		Logf(format, log...)
	}
}

func Fatalln(log ...interface{}) {
	if isLogAvailable(FATAL) {
		fmt.Print(concat("", BackgroundRed.String(), FATAL.String(), resetAll))
		Logln(log...)
	}
}

func Error(log ...interface{}) {
	if isLogAvailable(ERROR) {
		fmt.Print(concat("", BackgroundMagenta.String(), ERROR.String(), resetAll))
		Log(log...)
	}
}

func Errorf(format string, log ...interface{}) {
	if isLogAvailable(ERROR) {
		fmt.Print(concat("", BackgroundMagenta.String(), ERROR.String(), resetAll))
		Logf(format, log...)
	}
}

func Errorln(log ...interface{}) {
	if isLogAvailable(ERROR) {
		fmt.Print(concat("", BackgroundMagenta.String(), ERROR.String(), resetAll))
		Logln(log...)
	}
}

func Warn(log ...interface{}) {
	if isLogAvailable(WARN) {
		fmt.Print(concat("", BackgroundYellow.String(), WARN.String(), resetAll, " "))
		Log(log...)
	}
}

func Warnf(format string, log ...interface{}) {
	if isLogAvailable(WARN) {
		fmt.Print(concat("", BackgroundYellow.String(), WARN.String(), resetAll, " "))
		Logf(format, log...)
	}
}

func Warnln(log ...interface{}) {
	if isLogAvailable(WARN) {
		fmt.Print(concat("", BackgroundYellow.String(), WARN.String(), resetAll, " "))
		Logln(log...)
	}
}

func Info(log ...interface{}) {
	if isLogAvailable(INFO) {
		fmt.Print(concat("", BackgroundGreen.String(), INFO.String(), resetAll, " "))
		Log(log...)
	}
}

func Infof(format string, log ...interface{}) {
	if isLogAvailable(INFO) {
		fmt.Print(concat("", BackgroundGreen.String(), INFO.String(), resetAll, " "))
		Logf(format, log...)
	}
}

func Infoln(log ...interface{}) {
	if isLogAvailable(INFO) {
		fmt.Print(concat("", BackgroundGreen.String(), INFO.String(), resetAll, " "))
		Logln(log...)
	}
}

func Debug(log ...interface{}) {
	if isLogAvailable(DEBUG) {
		fmt.Print(concat("", BackgroundCyan.String(), DEBUG.String(), resetAll))
		Log(log...)
	}
}

func Debugf(format string, log ...interface{}) {
	if isLogAvailable(DEBUG) {
		fmt.Print(concat("", BackgroundCyan.String(), DEBUG.String(), resetAll))
		Logf(format, log...)
	}
}

func Debugln(log ...interface{}) {
	if isLogAvailable(DEBUG) {
		fmt.Print(concat("", BackgroundCyan.String(), DEBUG.String(), resetAll))
		Logln(log...)
	}
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
