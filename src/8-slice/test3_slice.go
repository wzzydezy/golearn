package main

import "fmt"

func main(){
	//slice1:= []int{1,2,3}
	//var slice1 []int
	//slice1 = make([]int, 3) //make开辟3个空间
	//var slice1 []int = make([]int, 3)
	//var slice1 = make([]int, 3)
    slice1:= make([]int, 3)
	slice1[0] = 100
	fmt.Printf("%d,%v",len(slice1), slice1)

	//判断slice是否为0
	if slice1 ==nil{
		fmt.Println("空切片")
	} else{
		fmt.Println("有空间")
	}
}