package main

import "fmt"


func calc(index string, a,b int) int{
	ret:= a+b
	fmt.Println(index,a,b,ret)
	return ret

}
//第一步：返回值赋值
//第二步：defer
//第三步：真正的RET返回
//函数中如果有defer ，那么执行的时机在第一步和第二步之间
func main(){
	a := 1
	b := 2
	defer calc("1",a,calc("10",a,b))      //
	a = 0
	defer calc("2",a,calc("20",a,b))
	b = 1


/*执行步骤：
1. a := 1
2. b := 2
3. defer calc("1",1,calc("10",1,2))
4. 先把函数内的返回：calc("10",1,2)           			//10 1 2 3
5. a = 0
6 .defer calc("2",0,calc("20",0,2))
7. 先把函数内的返回：calc("20",0,2)						//20 0 2 2
8. b = 1
9. defer calc("2",0, 2)					//2 0 2 2
10.defer calc("1",1, 3)					//1 1 3 4
*/

}
