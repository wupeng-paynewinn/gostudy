package mylogger

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

/*
日志库：
1、支持往不同地方输出日志
2、日志分级别
	1、DEBUG
	2、TRACE
	3、WARNING
	4、INFO
	5、ERROR
	6、FATAL
3、日志要支持开关控制
4、完整的记录日志要包含有时间、行号、文件名、日志级别、日志信息
5、日志文件要切割
	1、按文件大小切割
		1、每次记录日志之前判断一下当前写的这个文件的文件大小
	2、按日期切割
*/

//往终端写日志相关内容

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger interface {
	DEBUG(format string, a ...interface{})
	INFO(format string, a ...interface{})
	TRACE(format string, a ...interface{})
	WARNING(format string, a ...interface{})
	ERROR(format string, a ...interface{})
	FATAL(format string, a ...interface{})
}

// Logger 日志构造体
type ConsoleLogger struct {
	Level LogLevel
}
type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存的路径
	fileName    string //日志文件保存的名称
	maxFileSize int64
	fileObj     *os.File
	errFileObj  *os.File
	logChan     chan *LogMsg
}
type LogMsg struct {
	Level     LogLevel
	Msg       string
	funcName  string
	FileName  string
	Timestamp string
	Line      int
}

func parseLogLevel(level string) (LogLevel, error) {
	s := strings.ToUpper(level)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// NewLog 构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

//获取日志级别
func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

//获取行号、函数名、文件名
func getINFO(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip) //  1:往上再找一层  2:往上找两层
	if !ok {
		fmt.Println("runtime.caller() failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return

}
