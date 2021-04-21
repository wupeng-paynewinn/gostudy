package main

import (
"fmt"
"unicode"
)


func main(){
	//P15作业 ：查找字符串中的中文数量
	str := "135这是个一包含汉字和英文的字符串,skjdlkajkl"
	//count := len([]byte(str))
	//fmt.Println(count)
	hanNumber := 0
	digNumber := 0
	for _,v :=range str{
		if unicode.Is(unicode.Han,v){
			//fmt.Println(v)
			hanNumber++
		}
		if unicode.Is(unicode.Number,v){
			//fmt.Println(v)
			digNumber++
		}
	}

	// 常见的占位符：%c:字符，%d
	fmt.Printf("汉字的个数：%v\n",hanNumber)
	fmt.Printf("数字的个数：%v\n",digNumber)


}
