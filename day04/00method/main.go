package main

import "fmt"

//方法是有接收者的函数，接收者指的是哪个类型的变量可以调用这个函数


type person struct{
	name string
	age int
}

//只有person类型的接收者才能调用dream()这个方法
func (p *person) dream(d string){
	fmt.Printf("%s的梦想是：%s\n",p.name,d)
}

//指针接收者
//1、需要修改结构体变量的值时，要使用指针接收者
//2、结构体本身比较大，拷贝的内存开销比较大，也要使用指针接收者
//3、保持一致性：如果有一个方法使用了指针接收者，其他方法也统一要使用指针接收者
func (p *person) guonian(){
	p.age++
}
func main()  {
	p1 := person{name: "香蕉",age: 18}
	p2 := person{name: "爬山",age: 50}
	p1.dream("做个咸鱼")
	p2.dream("到山顶")
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
}