package logplus

import (
	"io"
	"os"
	"time"
)

var (
	lp *Logger
)

func init() {
	lp = NewLogger(os.Stdout, time.RFC3339, INFO, new(TextFormatter))
}

func SetTimeFormat(tf string) {
	lp.setTimeFormat(tf)
}

func SetLevel(level LogLevel) {
	lp.setLevel(level)
}

func SetOutput(out io.Writer) {
	lp.setOutput(out)
}

func SetFormatter(f Formatter) {
	lp.setFormatter(f)
}

func Log(log ...interface{}) {
	lp.print(log...)
}

func Logf(format string, log ...interface{}) {
	lp.printf(format, log...)
}

func Logln(log ...interface{}) {
	lp.println(log...)
}

func Colored(color Color, log ...interface{}) {
	lp.coloredPrint(color, log...)
}

func Coloredf(color Color, format string, log ...interface{}) {
	lp.coloredPrintf(color, format, log...)
}

func Coloredln(color Color, log ...interface{}) {
	lp.coloredPrintln(color, log...)
}

func Panic(log ...interface{}) {
	lp.leveledPrint(PANIC, log...)
}

func Panicf(format string, log ...interface{}) {
	lp.leveledPrintf(PANIC, format, log...)
}

func Panicln(log ...interface{}) {
	lp.leveledPrintln(PANIC, log...)
}

func Fatal(log ...interface{}) {
	lp.leveledPrint(FATAL, log...)
}

func Fatalf(format string, log ...interface{}) {
	lp.leveledPrintf(FATAL, format, log...)
}

func Fatalln(log ...interface{}) {
	lp.leveledPrintln(FATAL, log...)
}

func Error(log ...interface{}) {
	lp.leveledPrint(ERROR, log...)
}

func Errorf(format string, log ...interface{}) {
	lp.leveledPrintf(ERROR, format, log...)
}

func Errorln(log ...interface{}) {
	lp.leveledPrintln(ERROR, log...)
}

func Warn(log ...interface{}) {
	lp.leveledPrint(WARN, log...)
}

func Warnf(format string, log ...interface{}) {
	lp.leveledPrintf(WARN, format, log...)
}

func Warnln(log ...interface{}) {
	lp.leveledPrintln(WARN, log...)
}

func Info(log ...interface{}) {
	lp.leveledPrint(INFO, log...)
}

func Infof(format string, log ...interface{}) {
	lp.leveledPrintf(INFO, format, log...)
}

func Infoln(log ...interface{}) {
	lp.leveledPrintln(INFO, log...)
}

func Debug(log ...interface{}) {
	lp.leveledPrint(DEBUG, log...)
}

func Debugf(format string, log ...interface{}) {
	lp.leveledPrintf(DEBUG, format, log...)
}

func Debugln(log ...interface{}) {
	lp.leveledPrintln(DEBUG, log...)
}
