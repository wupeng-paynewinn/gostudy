package main

import (
	"fmt"
	"path"
	"runtime"
)

func getINFO() {
	pc, file, line, ok := runtime.Caller(0) //  1:往上再找一层 f1() 23行   2:往上找两层 main（） 27
	if !ok {
		fmt.Println("runtime.caller() failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file) // D:/GoProject/src/gostudy/day06/05runtime_demo/main.go
	fmt.Println(path.Base(file))
	fmt.Println(line) //9
}

func f1() {
	getINFO()
}

func main() {
	f1()
}
