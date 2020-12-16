package main

import (
	"fmt"
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)  //100000000000000000000
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main(){
	fmt.Println("Hello world!!")
	fmt.Println("Hello world!!")
	fmt.Println("Hello world!!")

	fmt.Println("KB:%S",KB)
	fmt.Println("MB:%S",MB)
	fmt.Println("GB:%S",GB)
	fmt.Println("TB:%S",TB)
	fmt.Println("PB:%S",PB)

}
