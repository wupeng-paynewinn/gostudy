package main

import "fmt"

func f1(a [3]int) (b [3]int){
	a[1] = 100
	b = a
	return b
}

func main(){
	//求a1数组的和
	//...标识位置的元素长度
	//a1 := [30]int  //30表示元素的个数（数组的长度） int表示元素的类型，数组的长度属于数组类型的一部分


	//数组的初始化
	//数组的第一种申明方法
	var a2 [8]int   //声明
	a2 = [8]int{1,2,4,5,6}		//初始化
	fmt.Println(a2)

	//数组的第三种声明方法：根据索引创建值
	a3 := [...]int{1:100,3:200}   //索引1的位置，值为100，索引99的位置，值为200
	fmt.Println(a3)

	//数组的第二种声明方法
	a1 	:=	[...]int{1,2,3,6,7,8}   //声明数组a1并初始化数组
	sum := 0
	//二维数组
	//var a4 [3][2]int   //[[1,2],[3,4].[4,5]]

	var a4 = [3][2]int{    //3的地方可以换为...
		[2]int{1,2},
		[2]int{3,4},
		[2]int{4,5},
	}
	fmt.Printf("a4:%v\n",a4)

	//数组是值类型
	x :=[3]int{1,2,3}
	y := x    //把x拷贝了一份给y
	y[1] = 200   //修改的是y的数据

	fmt.Printf("初始化x:%v\n" ,x)   //[1,2,3] //
	f1(x)
	fmt.Printf("调用方法后x:%v\n",x)   //[1 2 3]
	fmt.Printf("f1:%v\n" ,f1(x))  //[1 100 3]

	//求和
	for _,v := range a1{
		sum = sum +v
	}
	fmt.Printf("sum: %d\n",sum)


	//	求a1数组里和为7的两个元素的下标
	for i:=0;i<len(a1);i++{
		//fmt.Print(a1)
		for j:=i+1;j<len(a1);j++{
			if a1[i]+a1[j]==7{
				fmt.Printf("和为7的下标：(%d ,%d)\n",i,j)
			}
		}
	}


}


