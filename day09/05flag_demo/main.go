package main

import (
	"flag"
	"fmt"
	_ "os"
	"time"
)

// os.Args 获取命令行参数

func main() {
	// 第一种定义命令行flag参数的方法
	//name := flag.String("name","王也","输入姓名")
	//age := flag.Int("age",30,"输入年龄")
	//married := flag.Bool("married",false,"是否结婚")
	//mTime := flag.Duration("mTime",time.Second,"输入结婚时长")
	//
	//flag.Parse()
	//
	//fmt.Println(*name)
	//fmt.Println(*age)
	//fmt.Println(*married)
	//fmt.Println(*mTime)

	// 第二种定义命令行flag参数的方法
	var (
		name    string
		age     int
		married bool
		mTime   time.Duration
	)

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&mTime, "d", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(mTime)

	//运行命令：05flag_demo.exe -name 娜扎 --age 28 -married=false -mTime=1h30m a b c
	fmt.Println(flag.Args())  ////返回命令行参数后的其他参数，以[]string类型
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
