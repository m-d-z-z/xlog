package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

const callDepth = 3

type Level string

const (
	LvDebug   Level = "Debug"
	LvInfo    Level = "Info"
	LvWarning Level = "Warning"
	LvError   Level = "Error"
	LvPanic   Level = "PANIC"
	LvFatal   Level = "FATAL ERROR"
)

func SetLogLevel (level Level) {
	logWeight = getWeight(level)
}

var enableFilename = false

func EnableFilename() {
	enableFilename = true
}

var logWeight = 0

func getWeight(tag Level) int {
	switch tag {
	case LvDebug:
		return 1
	case LvInfo:
		return 2
	case LvWarning:
		return 3
	case LvError:
		return 4
	case LvPanic:
		return 5
	case LvFatal:
		return 6
	default:
	}
	return 999999
}

func getFile() string {
	if !enableFilename {
		return "-"
	}
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	}
	count := 0
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			count++
			if count == 2 {
				break
			}
		}
	}
	return fmt.Sprintf("%v:%d", short, line)
}

func logPrintf(tag Level, format string, v... interface{})  {
	if getWeight(tag) < logWeight {
		return
	}
	_ = log.Output(callDepth, fmt.Sprintf("%-7s %s %s", fmt.Sprintf("[%s]", tag), getFile(), fmt.Sprintf(format, v...)))
}

func logPrintln(tag Level, v... interface{})  {
	if getWeight(tag) < logWeight {
		return
	}
	_ = log.Output(callDepth, fmt.Sprintf("%-7s %s %s", fmt.Sprintf("[%s]", tag), getFile(), fmt.Sprint(v...)))
}

func Tag(tag string, v... interface{})  {
	logPrintln(Level(tag), v...)
}

func Tagf(tag string, format string, v... interface{})  {
	logPrintf(Level(tag), format, v...)
}

func Debug(v... interface{})  {
	logPrintln(LvDebug, v...)
}

func Info(v... interface{}) {
	logPrintln(LvInfo, v...)
}

func Warning(v... interface{}) {
	logPrintln(LvWarning, v...)
}

func Error(v... interface{}) {
	logPrintln(LvError, v...)
}

func Panic(v... interface{}) {
	logPrintln(LvPanic, v...)
}

func Fatal(v... interface{}) {
	logPrintln(LvFatal, v...)
	os.Exit(1)
}

func Debugf(format string, v... interface{})  {
	logPrintf(LvDebug, format, v...)
}

func Infof(format string, v... interface{}) {
	logPrintf(LvInfo, format, v...)
}

func Warningf(format string, v... interface{}) {
	logPrintf(LvWarning, format, v...)
}

func Errorf(format string, v... interface{}) {
	logPrintf(LvError, format, v...)
}

func Panicf(format string, v... interface{}) {
	logPrintf(LvPanic, format, v...)
}

func Fatalf(format string, v... interface{}) {
	logPrintf(LvFatal, format, v...)
	os.Exit(1)
}