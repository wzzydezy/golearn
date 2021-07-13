package main

import "fmt"

type Reader interface{
	ReadBook()
}

type Writer interface{
	WriteBook()
}

type Book struct {

}

//一个类实现多个接口
func (this *Book) ReadBook(){
	fmt.Println("R")
}

func (this *Book) WriteBook(){
	fmt.Println("W")
}

func main(){
	b := &Book{}


    //接口值pair<type,value>为空
	var r Reader

	r = b 

	r.ReadBook()

	var w Writer

	w = r.(Writer) //类型转换

	w.WriteBook()
}


