package main

import "fmt"

func main(){
	defer fmt.Println("main end1") //先入栈后执行
	defer fmt.Println("main end2") //后入栈先执行
	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}