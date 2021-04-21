package main

import "fmt"


type address struct {
	province string
	city string
}
type  workPlace struct {
	province string
	city string
}

type person struct {
	name string
	age int
	address
	workPlace
	//匿名嵌套 address  即：address:address

}
type company struct {
	name string
	addr address   //嵌套
}
func main() {
	p := person{
		name: "AirJordon",
		age: 18,
		address: address{
			province: "海南",
			city: "三亚",
		},
	}
	fmt.Println(p)
	fmt.Println(p.name,p.address.city)
	//fmt.Println(p.name,p.city)   //先去结构体里找，找不到则取匿名结构体里找
	fmt.Println(p.name,p.workPlace.city)
}
