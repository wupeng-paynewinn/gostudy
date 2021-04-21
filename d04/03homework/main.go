package main

import (
	"fmt"
	"os"
)

//学生管理系统

type  student struct {
	name string
	id int64
}

type studentMgr struct {
	allStudent map[int64]student
}
//查看学生
func (s studentMgr) showStudent()  {
	for _,stu := range s.allStudent{
		fmt.Printf("学号：%d 姓名：%s\n",stu.id,stu.name)
	}
}
//增加学生
func (s studentMgr) addStudent()  {
	var (
		stuID int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	//根据用户输入创建结构体对象
	newStu := student{
		id: stuID,
		name: stuName,
	}
	//2、把新的学生放到s.allStudent这个map中
	s.allStudent[newStu.id]=newStu
	fmt.Println("添加成功！")
}
//修改学生
func (s studentMgr) editStudent()  {
	//1、获取用户输入的学号
	var stuID int64
	fmt.Print("输入学号：")
	fmt.Scanln(&stuID)

	//2、展示学号对应的学生信息，如果没有，提示没有
	stuObj,ok := s.allStudent[stuID]
	if !ok{
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学号：%d,姓名：%s\n",stuObj.id,stuObj.name)
	//3、请输入修改后的学生名字
	var newName string
	fmt.Print("请输入要修改为什么姓名：")
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuID] = stuObj  //g更新map中的学生
	//4、更新学生的姓名


}
//删除学生
func (s studentMgr) delStudent()  {
	var stuID int64
	fmt.Print("请输入要删除学生的学号：")
	fmt.Scanln(&stuID)
	_,ok := s.allStudent[stuID]
	if !ok{
		fmt.Println("查无此人！")
		return
	}
	delete(s.allStudent,stuID)
	fmt.Println("删除成功！")
}

func showMenu(){
	fmt.Println(`
===Welcome sms！！===
	1、查看学生
	2、添加
	3、编辑
	4、删除
	5、推出
================`)
}
var smr studentMgr


func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student,100),
	}
	for {
		showMenu()
		//等待用户输入
		fmt.Print("请输入操作菜单的序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是：",choice)
		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.delStudent()
		case 5:
			os.Exit(1)
		}
	}
}
