package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())  //初始化随机数种子
	var scoreMap = make(map[string]int,200)

	for i:=1;i<100;i++{
		key := fmt.Sprintf("student_%02d",i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	//取出所有的map中的key存入切片keys
	var keys = make([]string,0,200)
	for key := range scoreMap{
		keys = append(keys,key)
	//对切片进行排序
	sort.Strings(keys)
	}

}
