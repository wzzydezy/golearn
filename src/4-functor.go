package main

import(
	"fmt"
)

func foo1(a string, b int) int{
	fmt.Println(a)
	fmt.Println(b)
	c:=100
	return c
}
//返回多个返回值，匿名
func foo2(a string, b int) (int, int){
	fmt.Println(a)
	fmt.Println(b)

	return 666,777
}
//返回多个返回值，有名
func foo3(a string, b int) (r1 int, r2 int){
	r1 = 1000
	r2 = 2000

	//若r1 r2不符值，默认初始化的值为0，有名形参
	//r1,r2的作用于空间就是函数空间
	return
}

func foo4(a string, b int) (r1, r2 int){
	r1 = 3000
	r2 = 4000
	return
}

func main(){
	c:=foo1("aaa",12)
	fmt.Println(c)

	ret1, ret2 := foo2("haha", 999)
	fmt.Println(ret1)
	fmt.Println(ret2)

	ret3, ret4 := foo3("haha", 999)
	fmt.Println(ret3)
	fmt.Println(ret4)

	ret5, ret6 := foo4("haha", 999)
	fmt.Println(ret5)
	fmt.Println(ret6)
}