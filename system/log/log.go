package log

import (
	"fmt"
	"github.com/wk331100/iFTY/config"
	"log"
	"os"
	"time"
)
const(
	TYPE_HOUR	= "hour"
	TYPE_DAY 	= "day"
	TYPE_MONTH	= "month"
	TYPE_YEAR	= "year"

	LEVEL_DEBUG = "debug"
	LEVEL_INFO 	= "info"
	LEVEL_WARNING = "warning"
	LEVEL_ERROR	= "error"
	LEVEL_FATAL = "fatal"

	DEFAULT_FILE_NAME = "errors"
)

func Debug(msg string) {
	writeLog(msg, LEVEL_DEBUG)
}

func Info(msg string) {
	writeLog(msg, LEVEL_INFO)
}

func Warning(msg string) {
	writeLog(msg, LEVEL_WARNING)
}

func Error(msg string) {
	writeLog(msg, LEVEL_ERROR)
}

func Fatal(msg string) {
	writeLog(msg, LEVEL_FATAL)
}


//初始化日志
func Init()  {
	appConfig := config.AppConfig
	logPath := appConfig["logPath"]

	//创建日志目录
	if err := os.MkdirAll(logPath.(string), 0766); err != nil {
		panic(err)
	}
}

//日志文件格式
func fileName(logType string, now time.Time) string {
	fileName := ""
	switch logType {
	case TYPE_HOUR:
		fileName = now.Format("15")
	case TYPE_DAY:
		fileName = now.Format(DEFAULT_FILE_NAME)
	case TYPE_MONTH:
		fileName = now.Format("2006_01")
	case TYPE_YEAR:
		fileName = now.Format("2006")
	default:
		fileName = now.Format(DEFAULT_FILE_NAME)
	}
	return  fileName + ".log"
}

func pathName(logPath, logType string, now time.Time) string {
	if logType == TYPE_HOUR{
		logPath += "/" + now.Format("20060102")
		if err := os.MkdirAll(logPath, 0766); err != nil {
			panic(err)
		}
	}
	return logPath
}


//写入日志
func writeLog(msg,level string)  {
	appConfig := config.AppConfig
	cstSh, _ := time.LoadLocation(appConfig["timeLocation"].(string))
	now := time.Now().In(cstSh)
	logPath := pathName(appConfig["logPath"].(string),appConfig["logType"].(string),now)

	fileName := fileName(appConfig["logType"].(string), now)

	fileFullPath := logPath +  "/" + fileName
	fmt.Println(fileFullPath)
	file, err := os.OpenFile(fileFullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND,  0766)
	defer file.Close()
	if err != nil {
		log.Println(err.Error())
	}

	msg = fmt.Sprintf("[%s] [%s] ", now.Format("2006-01-02 15:04:05"), level) + msg + "\n"
	fmt.Println(msg)
	_, err = file.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
}