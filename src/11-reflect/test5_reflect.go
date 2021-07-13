package main

import(
	"fmt"
	"reflect"
)

type User struct{
	Id int
	Name string
	Age int
}

func (this User) Call(){
	fmt.Println("this")
}

func main(){
	user := User{1,"ace",18}
	
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}){
	//获取type
	inputType := reflect.TypeOf(input)
	fmt.Println(inputType.Name())
	//获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println(inputValue)

	//通过Type获取字段名字
	//1.通过Type得到NumFiled，进行遍历
	//2.get every field, data type
	//3.get corrsponding value
	for i := 0; i< inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Println(field.Name, field.Type, value)
		fmt.Println(inputType.NumMethod())
	}

	for i := 0; i< inputType.NumMethod(); i++{
		m := inputType.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}