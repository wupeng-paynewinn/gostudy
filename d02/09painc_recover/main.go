package main

import "fmt"

func a()  {
	fmt.Println("a")
}


func b()  {
	//刚刚打开数据库连接

	//recover()必须搭配defer使用
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	if recover() != nil{
		panic("error:502")
	}else {
		fmt.Println("b")
	}

}


func c()  {
	fmt.Println("c")
}


func main()  {
	a()
	b()
	c()
}