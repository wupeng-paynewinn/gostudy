package main

import "fmt"

// P16课后作业：打印乘法口诀表  (梯形)
func main(){
	for i :=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		//fmt.Printf("%d\t",i)
		fmt.Printf("\n")
	}
}
