package main

import "fmt"

func addr(x int) func(int) int{
	return func(y int) int {
		x += y
		return x
	}
}

func main(){
	f := addr(10)(20)
	b := addr(1)(8)
	fmt.Println(f)
	fmt.Println(b)
}
