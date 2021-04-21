package main

import (
	"fmt"
	"strings"
)

func makeSuffixFuc(suffix string) func(string) string {
	return func(name string) string{
		if !strings.HasSuffix(name,suffix){
			return name + suffix
		}
		return suffix
	}
}

func main(){
 	jpgFunc := makeSuffixFuc(".jpg")
 	txtFunc := makeSuffixFuc(".txt")
 	fmt.Println(jpgFunc("test"))
 	fmt.Println(txtFunc("test"))
}
