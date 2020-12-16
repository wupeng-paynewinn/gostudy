package main

import (
	"fmt"
)
import "encoding/json"

//结构体与json

//1、序列化：把Go语言中的结构体变量 --> json格式的字符串
//2、反序列化： 把json格式的字符串 --> go语言中能够识别的结构体变量


//结构体内部的字段首字母要大写，不大写别人访问不到

type person struct{
	Name string  `json:"name" db:"name" ini:"name"`  //在用json时候，字段名为name小写。
	Age int `json:"age" `
}

func main() {
	p1 := person{
		Name: "萌新",
		Age: 8,
	}
	fmt.Println(p1)

	//序列化
	b ,err :=json.Marshal(p1)
	if err != nil{
		fmt.Printf("marshal failed ,err:%v",err)
		return
	}
	fmt.Printf("%v\n",string(b))

	//反序列化
	str := `{"name":"爱沙可","age":18}`
	var p2 person
	//反序列化时要传递指针
	err = json.Unmarshal([]byte(str), &p2) //传指针
	if  err !=nil{
		fmt.Printf("unmarshal failed:%v\n",err)
	}
	fmt.Printf("%#v\n",p2)
}