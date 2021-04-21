package main

import "fmt"

type  person struct {
	name,gender string
}

func f(x person)  {
	x.gender = "女"					//修改的是副本的gender
}

func f2(x *person)  {
	(*x).gender = "女"				//根据内存地址找到原位置
}
func main()  {
	var p person
	p.name ="乘客"
	p.gender="男"
	f(p)
	fmt.Println(p.gender)   		//男
	f2(&p)
	fmt.Println(p.gender)	 		 //女


	var p2 =new(person)
	(*p2).name = "计算机"
	p2.gender = "保密"
	fmt.Printf("%T\n",p2)
	fmt.Printf("%p\n",p2)
	fmt.Printf("%p\n",&p2)


	//2结构体指针
	//2.1 key=value初始化
	//常用的：
	var p3 = &person{
		name:"元素",
		gender: "男",
	}
	fmt.Printf("%#v\n",p3)

	//2.2使用列表的形式初始化,值的顺序要和结构体定义时字段的顺序一致
	p4 := &person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n",p4)
}
