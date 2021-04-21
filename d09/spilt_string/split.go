package spilt_string

import (
	"fmt"
	"strings"
)

// 写一个有bug的切割函数，用于单元测试
func Split(str, sep string) []string {
	if strings.Contains(str, sep) {
		//性能优化
		//var ret []string
		var ret = make([]string, 0, strings.Count(str, sep)+1)
		index := strings.Index(str, sep)
		for index >= 0 {
			ret = append(ret, str[:index])
			str = str[index+len(sep):]
			index = strings.Index(str, sep)
		}
		ret = append(ret, str)
		//ret := strings.Split(str,sep)
		return ret
	} else {
		null := []string{""}
		fmt.Printf("seq '%v' not in str '%v'\n", sep, str)
		return null
	}
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
