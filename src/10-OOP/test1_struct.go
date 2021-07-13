package main

import "fmt"

//声明一种新的数据类型 myint， 是int的一种类型别名
/*
type myint int

func main(){
	var a myint = 10
	fmt.Println(a)
}
*/

type Book struct{
	title string
	auth string
}

func changebook(book Book){
	//传递一个副本，值传递
	book.auth = "666"
}

func changebook2(book *Book){
	book.auth ="777"
}

func main(){
	var book1 Book
	book1.title = "Golang"
	book1.auth = "zhang3"
    changebook(book1)
	fmt.Println(book1)
	changebook2(&book1)
	fmt.Println(book1)
}