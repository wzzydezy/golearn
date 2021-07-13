package main

import "fmt"

type Human struct{
	name string
	sex string
}

func (this *Human) Eat(){
	fmt.Println("Eat")
}

func (this *Human) Walk(){
	fmt.Println("Walk")
}

type Superman struct{
	Human //表示Superman继承了Human类的方法

	level int 
}

//重定义父类的方法Eat()
func (this *Superman) Eat(){
	fmt.Println("SEat")
}

//子类的新方法
func (this *Superman) Fly(){
	fmt.Println("SFly")
}

func (this *Superman) Print(){
	fmt.Println(this.name, this.sex, this.level)
}

func main(){
	h := Human{"zhang3","female"}
	h.Eat()
	h.Walk()

	//定义一个子类对象
	//s:= Superman{Human{"zhang4","female"}, 80}
	var s Superman
	s.name = "li4"
	s.sex = "male"
	s.level = 88

	s.Walk() //父类的方法
	s.Eat()  //子类的方法
	s.Fly()  //子类的方法
	s.Print()
}