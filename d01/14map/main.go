package main

import (
	"fmt"
)

func main(){
	//map的定义：var name map[key的type]：value的type
	var m1 map[string]int    //没有在内存中开辟空间
	m1 = make(map[string]int,10)   //为map申请10内存
	m1["年领"] = 30
	m1["存量"] = 20
	fmt.Println(m1)
	fmt.Println("不存在的key：",m1["月卡"])
	value,ok := m1["月卡"]
	//如果返回值是布尔值，通常用ok去接收它
	if !ok {
		fmt.Println("无此key")
	}else {
		fmt.Printf("key:%v",value)
	}

	//只遍历key
	for k := range m1{
		fmt.Println(k)
	}

	fmt.Printf("只遍历value:\n")
	//只遍历value
	for _,v :=range m1{
		fmt.Println(v)
	}

	fmt.Printf("遍历key，value:\n")
	//遍历key，value
	for k,v :=range m1{
		fmt.Println(k,v)
	}

	fmt.Printf("删除map:\n")
	//删除map
	delete(m1,"存量")
	fmt.Println(m1)
	delete(m1,"阿萨德") //delete 不存在的key，则不做任何操作

}