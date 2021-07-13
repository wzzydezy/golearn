package main

import "fmt"


//
//interface{} 空接口
//本质是一个指针
type Animal interface{
	Sleep()
	GetColor() string
	GetType() string
}

//定义一个具体的类
type Cat struct{
	color string //cat color
}

func (this *Cat) Sleep(){
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct{
	color string //cat color
}

func (this *Dog) Sleep(){
	fmt.Println("Dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal){
	animal.Sleep()
	fmt.Println(animal.GetColor())
	fmt.Println(animal.GetType())
}


func main(){
	var animal Animal
	animal = &Cat{"Green"}
	animal.Sleep()
	animal = &Dog{"Red"}
	animal.Sleep()

	cat := Cat{"G"}
	dog := Dog{"Y"}

	showAnimal(&cat)
	showAnimal(&dog)
}