package main

import "fmt"

//结构体是值类型
type person struct {
	name string
	age int
}

//构造函数：约定成俗以new开头
//返回的是结构体还是结构体指针:
//当结构体比较大的时候，尽量使用结构体指针，减少程序的内存开销
func newPerson(name string,age  int ) *person{
	return &person{
		name:name,
		age: age,
	}
}
func main() {
	p1 := newPerson("July",18)
	p2 := newPerson("Ada",78)
	fmt.Println(p1,p2)
	fmt.Printf("%T\n,%T\n",p1,p2)
}
