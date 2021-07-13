package main


import (
	"fmt"
	"reflect"
)
func reflectNum(arg interface{}){
	fmt.Println((reflect.TypeOf(arg)).Name())
	fmt.Println(reflect.ValueOf(arg))
}

type input struct{
	Id int
	Name string
}


func main(){
	var num float64 = 1.2345
	
	var in input
	reflectNum(num)
	reflectNum(&in)
}