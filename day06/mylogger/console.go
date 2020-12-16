package mylogger

import (
	"fmt"
	"time"
)

// Logger 日志构造体
type Logger struct {
}

// NewLog 构造函数
func NewLog() Logger {
	return Logger{}
}

//方法
func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Debug] %s \n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Info] %s \n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Warning] %s \n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Error] %s \n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Fatal] %s \n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Trace(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Trace] %s \n", now.Format("2006-01-02 15:04:05"), msg)
}
