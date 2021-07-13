package main

import "fmt"

func main(){
	//声明map1是一种map类型，key是string， value也是string
	var mymap1 map[string]string
    if mymap1 ==nil{
		fmt.Println("空map")
	}
	//在使用map前，需要使用make给map分配数据空间
	mymap1 = make(map[string]string, 10)
	
	mymap1["one"] = "jave"
	mymap1["two"] = "c++"
	mymap1["three"] = "python"

	fmt.Println(mymap1)

	mymap2 := make(map[int]string)
	mymap2[1] = "jave"
	mymap2[2] = "c++"
	mymap2[3] = "python"

	fmt.Println(mymap2)

	mymap3 := map[string]string{
		"one":"php",
		"two":"c++",
		"three":"go",
	}

	fmt.Println(mymap3)




}