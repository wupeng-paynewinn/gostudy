package main

import "fmt"

//匿名字段
//适用于字段比较少也比较简单的场景
type person struct {
	string
	int
}


func main() {
	p := person{
		"赵本山",
		19,
	}
	fmt.Println(p.int)
	fmt.Println(p.string)
}
