package calc

//包中标识符（变量名、函数名、结构体名、接口名、常量...）可见性 -->标识符大写表示对外可见
//大写的别的包才能访问，小写的表示私有的
func Add(x,y int) int  {
	return x+y
}

func Multi(x,y int) int {
	return x*y
}