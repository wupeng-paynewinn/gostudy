package main

import (
	"gostudy/day06/mylogger"
)

var log mylogger.Logger

//测试自己写的日志库
func main() {
	log = mylogger.NewConsoleLog("INFO")                                //终端日志实例
	log = mylogger.NewFileLogger("INFO", "./", "file.log", 1*1024*1024) //文件日志实例
	for {
		log.DEBUG("This is a DEBUG msg")
		log.TRACE("This is a TRACE msg")
		log.INFO("This is a INFO msg")
		log.WARNING("This is a WARNING msg")

		id := 10010
		name := "用户10010"
		reason := "这是一条报错信息"
		log.ERROR("This is a ERROR msg,id：%d,name:%s,err:%s", id, name, reason)

		log.FATAL("This is a FATAL msg")
		//time.Sleep(time.Second * 1)
	}
}
