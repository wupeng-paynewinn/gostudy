package main

import "fmt"

//多个包的init执行顺序：全局申明-->init()-->main()

//一个包里面只有一个init()
//一般用于做一些初始化的操作（初始化数据库连接等）

func init()  {
	fmt.Println("自动执行，没有参数，没有返回值，也不能手动调用")
}
