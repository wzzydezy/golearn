package main

import "fmt"

func printArray(myArray []int){
	//_表示匿名的变量
	//切片传递是引用传递,数组传递是值传递
	for _, value := range myArray{
       fmt.Println(value)
	}
}

func main(){
	myArray := []int{1,2,3,4}
	fmt.Printf("%T\n", myArray)

	printArray(myArray)
}