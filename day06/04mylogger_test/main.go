package main

import (
	"gostudy/day06/mylogger"
	"time"
)

//测试自己写的日志库
func main() {
	log := mylogger.NewLog()
	for {
		log.Debug("debug msg")
		log.Info("info msg")
		log.Error("error msg")
		log.Fatal("Fatal msg")
		log.Trace("Trace msg")
		log.Warning("warning msg")
		time.Sleep(time.Second * 5)
	}

}
