package main

import (
	"fmt"
	"os"
)

/*
	函数学生管理系统
	写一个函数能查看/新增/删除学生
 */

type student struct {
	id int64
	name string
}

var(
	allStudent  map[int64]*student
)

func showAllStudent()  {
	if len(allStudent) != 0{
		for k,v := range allStudent{
			fmt.Printf("id:%d , name:%s \n\n",k,v.name)
		}
	}else {
		fmt.Println("没人了呀，兄dei~")
	}

}


func newStu(id int64,name string)  *student{
	return &student{
		id:id,
		name: name,
	}
}


func addStudent()  {
	var (
		id int64
		name string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&name)
	newStudent :=newStu(id,name)
	allStudent[id] = newStudent
	fmt.Printf("学号：%d，姓名：%s  信息添加成功\n",id,name)
}


func deleteStudent()  {
	var(
		deleteId int64
	)
	fmt.Print("输入要删除的学号：")
	fmt.Scanln(&deleteId)
	//删除对应的键值对
	delete(allStudent,deleteId)
	fmt.Printf("学号：%d删除成功\n",deleteId)
}


func main() {
	/*
	1、打印菜单
	2、用户选择
	3、执行对应的函数
	 */
	allStudent = make(map[int64]*student,48)
	for {
		fmt.Println("Welcome to student system！")
		fmt.Println(`
				1、查看学生
				2、新增学生
				3、删除学生
				4、退出`)
		fmt.Println("输入你的选择：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d\n", choice)
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(-1) //退出
		default:
			fmt.Println("滚,没有~")
		}
	}
}
