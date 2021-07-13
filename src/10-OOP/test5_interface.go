package main

import "fmt"

//interface{}万能类型
func myfunc(arg interface{}){
	fmt.Println("myfunc is called")
	fmt.Println(arg)

	//interface{} 如何区分底层数据类型呢？

	//给interface{}提供“类型断言”的机制
	value, ok := arg.(string)
	if !ok{
		fmt.Println("T")
	}else{
		fmt.Println(value)
		fmt.Printf("%T",value)
	}


}

type Book struct{
	auth string
}

func main(){
	book:= Book{"Golong"}
	myfunc(book)
	myfunc(100)
	myfunc("abc")
}