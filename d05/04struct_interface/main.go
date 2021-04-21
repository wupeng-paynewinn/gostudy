package main

import "fmt"

//实现多个接口和接口嵌套
type animal interface {
	mover  //move()
	eater	//eat() //接口嵌套
}

type mover interface {
	move()
}
type eater interface {
	eat(s string)
}


type cat struct {
	name string
	feet int8
}

//cat实现了mover接口
func (c *cat)move()  {
	fmt.Println("猫动")
}
//cat实现了eater接口
func (c  *cat)eat(food string)  {
	fmt.Printf("吃%s!\n",food)
}

func main() {

}

