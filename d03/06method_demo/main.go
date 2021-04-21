package main

import "fmt"

//方法
//标识符：变量名 函数名 类型名 方法名
//Go语言中，如果标识符首字符是大写，表示对外部包可见、可调用（暴露的，公开的）

//对首字母大写的结构体要求有注释
//Dog 这是一个狗的结构体
type dog struct {
	name string
	age int
}

//构造函数
func newDog(name string,age int)  dog{
	return dog{
		name:name,
		age:age,
	}

}
//方法是作用于特定类型的函数
//接受者表示的是调用该方法的具体类型变量，多用类型变量首字母小写表示
//作为一个方法，必须要指定接收者：多了一个(d dog)，表示接收者
//接收者只能有一个
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~\n",d.name)
}

//使用值接收者
func (d dog) growUp()  {
	d.age++
}

//使用指针接收者：
//1 需要修改接收者的值；2 接收者拷贝代价大的时候；3 保证一致性（当一个方法使用了指针接收者，其他方法尽量使用指针接收者）
func (d *dog)  realGrowUp(){
	d.age++
}

func (d *dog) dream()  {
	fmt.Println("eat!!...more...")
}
func main() {
	d1 := newDog("花花",8)
	d1.wang()

	fmt.Println(d1.age)
	d1.growUp()    //值类型 不修改原来的值,拷贝
	fmt.Println(d1.age)

	d1.realGrowUp() //传地址，修改原来的值
	fmt.Println(d1.age)
}
