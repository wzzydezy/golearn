package main

import(
	"fmt"
)

//const定义枚举类型
const(
	//可以在const() 添加一个关键字iota，每行的iota都会累加1，第一行的iota默认值是0
	BEIJING = 10*iota //1
	SHANGHAI       //2
	SHENZHEN       //3
)

const(
	a,b = iota+1, iota+2
	c,d 
	e,f
	g,h = iota*2, iota*3
	i,k
)

func main(){
	//常量
	const length int = 10
	fmt.Println(a,b,c,d,e,f,g,h,i,k)
	//iota只能配合const()一起使用，iota只有在const进行累加效果
}