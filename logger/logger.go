package logger

import (
	"io"
	"log"
)

var (
	SUCCESS *log.Logger
	TRACE   *log.Logger
	INFO    *log.Logger
	WARN    *log.Logger
	ERROR   *log.Logger
	FATAL   *log.Logger
)

// Initialize Loggers
func InitLogger(
	successHandle io.Writer,
	traceHandle io.Writer,
	infoHandle io.Writer,
	warnHandle io.Writer,
	errorHandle io.Writer,
	fatalHandle io.Writer) {
	flag := log.Ldate | log.Ltime | log.Lshortfile

	SUCCESS = log.New(successHandle, "SUCCESS: ", flag)
	TRACE = log.New(traceHandle, "TRACE: ", flag)
	INFO = log.New(infoHandle, "INFO: ", flag)
	WARN = log.New(warnHandle, "WARN: ", flag)
	ERROR = log.New(errorHandle, "ERROR: ", flag)
	FATAL = log.New(fatalHandle, "FATAL: ", flag)
}
