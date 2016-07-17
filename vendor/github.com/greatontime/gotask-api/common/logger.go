package common

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	UNSPECIFIED Level = iota
	TRACE
	INFO
	WARNING
	ERROR
)

type Level int

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	isFlag bool) {
	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}
	Trace = log.New(traceHandle, "TRACE:", flag)
	Info = log.New(infoHandle, "INFO:", flag)
	Warning = log.New(warningHandle, "WARNING:", flag)
	Error = log.New(errorHandle, "ERROR:", flag)
}
func setLogLevel(level Level) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error openning log file : %s", err.Error())
	}
	switch level {
	case TRACE:
		initLog(f, f, f, f, true)
		return
	case INFO:
		initLog(ioutil.Discard, f, f, f, true)
		return
	case WARNING:
		initLog(ioutil.Discard, ioutil.Discard, f, f, true)
		return
	case ERROR:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, f, true)
		return
	default:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, false)
		f.Close()
		return
	}
}
