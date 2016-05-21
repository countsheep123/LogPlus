//	logplus.Logln("this is example")

package logplus

import (
	"fmt"
	"os"
	"time"
)

var (
	timeFormat = time.RFC3339
	logLevel   = INFO
)

func SetTimeFormat(tf string) {
	timeFormat = tf
}

func SetLevel(level LogLevel) {
	logLevel = level
}

func Log(log ...interface{}) {
	printCallInfo(3)

	fmt.Print(log...)
}

func Logf(format string, log ...interface{}) {
	printCallInfo(3)

	fmt.Printf(format, log...)
}

func Logln(log ...interface{}) {
	printCallInfo(3)

	fmt.Println(log...)
}

func Colored(color Color, log ...interface{}) {
	printCallInfo(3)

	s := fmt.Sprint(log...)
	fmt.Print(concat("", color.String(), s, resetAll))
}

func Coloredf(color Color, format string, log ...interface{}) {
	printCallInfo(3)

	s := fmt.Sprintf(format, log...)
	fmt.Print(concat("", color.String(), s, resetAll))
}

func Coloredln(color Color, log ...interface{}) {
	printCallInfo(3)

	s := fmt.Sprint(log...)
	fmt.Println(concat("", color.String(), s, resetAll))
}

func Panic(log ...interface{}) {
	if isLogAvailable(PANIC) {
		printLogInfo(BackgroundRed, PANIC)

		fmt.Print(log...)

		panic(fmt.Sprint(log...))
	}
}

func Panicf(format string, log ...interface{}) {
	if isLogAvailable(PANIC) {
		printLogInfo(BackgroundRed, PANIC)

		fmt.Printf(format, log...)

		panic(fmt.Sprintf(format, log...))
	}
}

func Panicln(log ...interface{}) {
	if isLogAvailable(PANIC) {
		printLogInfo(BackgroundRed, PANIC)

		fmt.Println(log...)

		panic(fmt.Sprintln(log...))
	}
}

func Fatal(log ...interface{}) {
	if isLogAvailable(FATAL) {
		printLogInfo(BackgroundRed, FATAL)

		fmt.Print(log...)

		os.Exit(1)
	}
}

func Fatalf(format string, log ...interface{}) {
	if isLogAvailable(FATAL) {
		printLogInfo(BackgroundRed, FATAL)

		fmt.Printf(format, log...)

		os.Exit(1)
	}
}

func Fatalln(log ...interface{}) {
	if isLogAvailable(FATAL) {
		printLogInfo(BackgroundRed, FATAL)

		fmt.Println(log...)

		os.Exit(1)
	}
}

func Error(log ...interface{}) {
	if isLogAvailable(ERROR) {
		printLogInfo(BackgroundMagenta, ERROR)

		fmt.Print(log...)
	}
}

func Errorf(format string, log ...interface{}) {
	if isLogAvailable(ERROR) {
		printLogInfo(BackgroundMagenta, ERROR)

		fmt.Printf(format, log...)
	}
}

func Errorln(log ...interface{}) {
	if isLogAvailable(ERROR) {
		printLogInfo(BackgroundMagenta, ERROR)

		fmt.Println(log...)
	}
}

func Warn(log ...interface{}) {
	if isLogAvailable(WARN) {
		printLogInfo(BackgroundYellow, WARN)

		fmt.Print(log...)
	}
}

func Warnf(format string, log ...interface{}) {
	if isLogAvailable(WARN) {
		printLogInfo(BackgroundYellow, WARN)

		fmt.Printf(format, log...)
	}
}

func Warnln(log ...interface{}) {
	if isLogAvailable(WARN) {
		printLogInfo(BackgroundYellow, WARN)

		fmt.Println(log...)
	}
}

func Info(log ...interface{}) {
	if isLogAvailable(INFO) {
		printLogInfo(BackgroundGreen, INFO)

		fmt.Print(log...)
	}
}

func Infof(format string, log ...interface{}) {
	if isLogAvailable(INFO) {
		printLogInfo(BackgroundGreen, INFO)

		fmt.Printf(format, log...)
	}
}

func Infoln(log ...interface{}) {
	if isLogAvailable(INFO) {
		printLogInfo(BackgroundGreen, INFO)

		fmt.Println(log...)
	}
}

func Debug(log ...interface{}) {
	if isLogAvailable(DEBUG) {
		printLogInfo(BackgroundCyan, DEBUG)

		fmt.Print(log...)
	}
}

func Debugf(format string, log ...interface{}) {
	if isLogAvailable(DEBUG) {
		printLogInfo(BackgroundCyan, DEBUG)

		fmt.Printf(format, log...)
	}
}

func Debugln(log ...interface{}) {
	if isLogAvailable(DEBUG) {
		printLogInfo(BackgroundCyan, DEBUG)

		fmt.Println(log...)
	}
}

const (
	resetAll = "\033[0m"
)

func printLogInfo(color Color, level LogLevel) {
	fmt.Print(concat("", color.String(), level.String(), resetAll), " ")
	printCallInfo(4)
}

func printCallInfo(depth int) {
	t := time.Now().Format(timeFormat)
	info := getCallInfo(depth)
	fmt.Printf("%s [%s] <%d> ", t, info, os.Getpid())
}

func isLogAvailable(level LogLevel) bool {
	if level <= logLevel {
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
