package main

import "fmt"

//如果类名首字母大写，表示其他包也能访问
type Hero struct{
	//如果类的属性首字母大写，表示该属性是public, 小写则是private
	Name string
	Ad int
	Level int
}

/*
func (this Hero) ShowName(){
	fmt.Println("Name = ",this.Name)
}

func (this Hero) GetName() string {
	return this.Name
}

func (this Hero) SetName(newName string){
	//this 是调用该方法的对象的一个拷贝
	this.Name = newName
}
*/

func (this *Hero) ShowName(){
	fmt.Println("Name = ",this.Name)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string){
	this.Name = newName
}

func main(){
	hero := Hero{Name:"zhange", Ad:3, Level:1}
	hero.ShowName()

	hero.SetName("li4")

	hero.ShowName()

	
}