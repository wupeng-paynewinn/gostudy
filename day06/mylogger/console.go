package mylogger

import (
	"fmt"
	"time"
)

func (c ConsoleLogger) enable(loglevel LogLevel) bool {
	return loglevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getINFO(3)
		fmt.Printf("[%s] [%s] [%s:%s:line%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}

}

//方法
func (c ConsoleLogger) DEBUG(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) TRACE(format string, a ...interface{}) {
	c.log(TRACE, format, a...)
}

func (c ConsoleLogger) INFO(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

func (c ConsoleLogger) WARNING(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) ERROR(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) FATAL(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
