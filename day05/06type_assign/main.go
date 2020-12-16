package main

import "fmt"

//类型断言：前提


//断言类型1：
func assign(a interface{})  {
	fmt.Printf("%T\n",a)
	str,ok := a.(string)
	if !ok{
		fmt.Println("猜错了")
	}else {
		fmt.Println("是字符串：",str)
	}

}

//类型断言2
func assign2(a interface{})  {
	fmt.Printf("%T\n",a)
	//fmt.Printf("#a{}\n")
	switch t := a.(type) {
	case string:
		fmt.Println("是字符串",t)
	case int:
		fmt.Println("是int：",t)
	case int64:
		fmt.Println("是int64",t)
	case bool:
		fmt.Println("是bool：",t)
	default:
		fmt.Println("无分类")
	}
}

func main()  {
	assign(100)
	assign2(int64(200))

}