package main

import "fmt"

// func swap1(a int, b int){//值传递改变形参，不改变实参
// 	
// }

func swap2(a *int, b *int){
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main(){
	var a int = 10
	var b int = 20
    swap2(&a, &b)
	fmt.Println(a,b)
}
