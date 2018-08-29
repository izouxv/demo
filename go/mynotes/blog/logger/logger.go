package logger

import (
	"os"
	"io"
	"log"
	"runtime"
	"strings"
	"fmt"
	"mynotes/blog/config"
)

var (
	debugLog   	*log.Logger // 记录debug日志
	infoLog    	*log.Logger // 记录info日志
	warnLog		*log.Logger // 记录warn日志
	errorLog	*log.Logger // 记录error日志
	debugFlag	=	"debug: "
	infoFlag	=	"info: "
	warnFlag	=	"warn: "
	errorFlag	=	"error: "
	split		=	"/"
	colon		=	":"
	space		=	" "
	unknownFile	=	"???"
	//errorOpenFile = "Failed to open log file:"
)

func InitLogger(logPath *config.LogPath) {
	if logPath == nil {
		logPath = &config.LogPath{}
	}
	if debugLog != nil || infoLog != nil || warnLog != nil || errorLog != nil {
		return
	}
	debugLog = initLogger(debugLog, logPath.Debug, debugFlag)
	infoLog = initLogger(infoLog, logPath.Info, infoFlag)
	warnLog = initLogger(warnLog, logPath.Warn, warnFlag)
	errorLog = initLogger(errorLog, logPath.Error, errorFlag)
}

func pathCaller() (string,int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = unknownFile
		line = 0
	}
	return file[strings.LastIndex(file,split)+1:],line
}

func initLogger(logger *log.Logger, loggerFile, loggerFlag string) *log.Logger {
	if loggerFile != "" {
		fmt.Println("aaa:",loggerFile)
		file, err := os.OpenFile(loggerFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		logger = log.New(io.MultiWriter(file,os.Stdout),loggerFlag,log.LstdFlags)
	} else {
		logger = log.New(os.Stdout,loggerFlag,log.LstdFlags)
	}
	return logger
}

func Debug(i... interface{}) {
	fileName, line := pathCaller()
	debugLog.Print(fileName+colon,line,space,fmt.Sprint(i...))
}

func Info(i... interface{}) {
	fileName, line := pathCaller()
	infoLog.Print(fileName+colon,line,space,fmt.Sprint(i...))
}

func Warn(i... interface{}) {
	fileName, line := pathCaller()
	warnLog.Print(fileName+colon,line,space,fmt.Sprint(i...))
}

func Error(i... interface{}) {
	fileName, line := pathCaller()
	errorLog.Print(fileName+colon,line,space,fmt.Sprint(i...))
}
