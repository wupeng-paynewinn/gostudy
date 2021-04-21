package main

import "fmt"

//结构体
type person struct {   //类型名称唯一
	name string			//字段名称唯一
	age int
	gender string
	hobby []string
}

func main()  {
	var a person
	a.name = "David"
	a.age = 19
	a.gender = "male"
	a.hobby = []string{"sleep","football","piano"}
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	fmt.Println(a.name)

	//匿名结构体:多用于临时场景
	var s struct{
		name  string
		age int
	}
	s.name = "ja"
	s.age = 18
	fmt.Println(s)
	fmt.Printf("%T\n",s)
	fmt.Println(s.name)



}
