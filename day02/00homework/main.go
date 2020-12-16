package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main(){
	s1 := "Hello卡了"
	var count int
	for _,v  := range s1{
		if unicode.Is(unicode.Han,v){
			count++
		}
	}
	fmt.Println(count)


	//判断how do you do 单词出现的次数
	//1 把字符串按照空格切割
	//2 遍历切片存在一个map
	s2 := "how do you do"
	s3 :=strings.Split(s2," ")
	m1 := make(map[string]int,10)
	for _,w := range s3{
		//如果map中不存在w这个key，key出现的次数+1
		if _,ok := m1[w];!ok{
			m1[w] = 1
		}else {  //如果存在w这个key，出现的次数+1
			m1[w]++
		}
	}
	//累计出现的次数
	for k,v :=range m1{
		fmt.Println(k,v)
	}


	//回文判断
	//字符串从左往右和从右往左读是一样的
	ss := "上海自来水来自海上"
	//思路：把字符串拿出来放到一个[]rune中
	r := make([]rune,0 ,len(ss))
	for _,i := range ss{
		r = append(r,i)
	}
	fmt.Println("[]rune:",r)
	for i := 0;i < len(r)/2;i++{
		if r[i] != r[len(r)-1-i]{
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
