package main

import (
	"fmt"
)

var (
	coins = 50
	user = []string{
		"Matthew","sarah","Augustus","Heidi","Emilie","Peter","Giana","Adriano","Aaron","Elizabeth",
	}
	distribution = make(map[string]int,len(user))
)
/*
50枚金币，需要分配给user们：
名字包含eE，分1枚
名字包含iI，分2枚
名字包含oO，分3枚
名字包含uU，分4枚

计算每个用户分到多少金币，以及剩下多少金币数？

 */
func dispatchCoin() (left int){
	/*
		1.依次拿到每个人的名字
		2.拿到名字根据规则分金币
		2.1 每个人分的存到distribution 中
		2.2 记录下剩余的金币数
		3。
	*/
	for _,name := range user{
		for _,c:= range name{
			switch c {
			case 'e' ,'E':
				distribution[name] ++
				coins --
			case 'i' ,'I':
				distribution[name] +=2
				coins -= 2
			case 'o' ,'O':
				distribution[name] +=3
				coins -= 3
			case 'u' ,'U':
				distribution[name] +=4
				coins -= 4
			}
		}
	}
	left = coins

	return
}

func main()  {
	left := dispatchCoin()
	fmt.Println("剩下：",left)
	for k,v := range distribution{
		fmt.Printf("%s:%d\n",k,v)
	}

}