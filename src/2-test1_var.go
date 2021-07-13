package main

import (
	"fmt"
)
//声明全局变量 方法123可以的
var gA int = 100
var gB = 200

//方法4声明全局变量
//:=只能用于函数体内
//gC:= 100报错


func main(){
	//1声明变量,默认值是0
	var a int
	fmt.Println("a = ",a)
	fmt.Printf("type of a = %T\n", a)
	//2
	var b int = 100
	fmt.Println("b = ",b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)
	//3省去数据类型
	var c = 100
	fmt.Println("c = ",c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)
	//4
	e := 100
	fmt.Println("e=",e)
	fmt.Printf("type e = %T\n",e)

	f := "abcd"
	fmt.Println("f=",f)
	fmt.Printf("type f = %T\n",f)

	g := 3.14
	fmt.Println("g=",g)
	fmt.Printf("type g = %T\n",g)
	
	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println(xx,yy)

	var kk, ll = 100, "abcd"
	fmt.Println(kk,ll)

	//多行多变量声明
	var(
		vv int = 100
		jj bool = true
	)
}