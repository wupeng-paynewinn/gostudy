package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"阿斯蒂","age":18}`
	var p person
	err := json.Unmarshal([]byte(str), &p)
	if err != nil {
		fmt.Printf("unmarshal failed:%v\n", err)
	}
	fmt.Println(p.Name, p.Age)
}
