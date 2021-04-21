package main

import (
	"fmt"

	//包名的导入是从$GOPATH下的src开始计算的，用/分隔：导入时从src后的路径写起就好了
	"study/day05/07calc"

	//Go语言中禁止循环导入包（A导入B，B导入C，C导入A）
)

func main()  {
	addNumber :=calc.Add(10,20)
	multiNumber :=calc.Multi(10,20)
	fmt.Println(addNumber)
	fmt.Println(multiNumber)
}
