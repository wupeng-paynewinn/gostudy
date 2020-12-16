package main

import "fmt"

//所有类型都是

//空接口作为函数参数

//interface 	：空接口关键字
//interface{} 	：空接口类型
func show(a interface{})  {
	fmt.Printf("type:%T value:%v\n",a,a)
}
func main() {
	m1 := make(map[string]interface{},16)
	m1["name"] = "蔡徐坤"
	m1["age"] = 8
	m1["married"] = true
	m1["hobby"] = [...]string{"唱","跳","rap"}
	fmt.Println(m1)

	show(0)
	show(false)
	show(nil)
	show(m1)
}
