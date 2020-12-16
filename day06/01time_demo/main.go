package main

import (
	"fmt"
	"time"
)

func f1() {
	now := time.Now()  //本地的时间
	//fmt.Println(now.Year(),now.Month(),now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour(),now.Minute(),now.Second())


	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) //纳秒

	//时间戳转换为时间格式
	t1 := time.Unix(1608022492,0)
	fmt.Println(t1)
	fmt.Println(t1.Year(),t1.Month(),t1.Day())

	//now + 24小时
	fmt.Println(now.Add(24*time.Hour))

	fmt.Println("---------------------------")
	//其他sub()相差；equal()相等，返回一个bool值；before() bool ；after bool;
	nextYear,err := time.Parse("2006-01-02","2020-12-25")
	if err != nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Println(nextYear.Sub(now))
	fmt.Println(now.Sub(nextYear))
	fmt.Println("---------------------------")

	//定时器
	//timer := time.Tick(time.Second)
	//for t := range timer{
	//	fmt.Println(t)
	//}

	//格式化时间
	//	2020-12-15
	fmt.Println(now.Format("2006-01-02"))
	//	2020/12/15 17:21:45
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//	2020/12/15 5:21:45 PM
	fmt.Println(now.Format("2006/01/02 3:04:05 PM"))
	//	2020/12/15 17:21:45.342
	fmt.Println(now.Format("2006/01/02 15:04:05.000"))


	//按照对应的格式解析字符串类型的时间
	timeObj,err := time.Parse("2006-01-02","2020-12-15")
	if err != nil{
		fmt.Println("err:",err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())

	fmt.Println("=======================")
	n :=5
	fmt.Println(time.Now())
	time.Sleep(time.Duration(n)*time.Second)
	fmt.Println(time.Now())
	fmt.Println("=======================")
}

func f2() {
	
}
func main() {
	f1()
}

