package main

import (
	"fmt"
	"strconv"
)

// strconv

func a2i() {
	//把字符串转换为数字类型
	str := "10000"
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)
	// Array to Int
	retInt, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	fmt.Printf("%#v %T\n", retInt, retInt)

}
func i2a() {
	// 把数字类型转换为字符串
	i := int32(2316)
	//ret2 :=string(i)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v %T\n", ret2, ret2)
	retArr := strconv.Itoa(int(i))
	fmt.Printf("%#v %T\n", retArr, retArr)
}
func a2b() {
	//字符串解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)
}

func a2f() {
	//字符串中解析出浮点型
	folatStr := "1.2345"
	folatValue, _ := strconv.ParseFloat(folatStr, 64)
	fmt.Printf("%#v %T\n", folatValue, folatValue)
}
func main() {
	a2i()
	i2a()
	a2b()
	a2f()
}
