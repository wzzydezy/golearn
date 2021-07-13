package main

import "fmt"

func printmap(citymap map[string]string){
	//map是引用传递
	for key, value := range citymap{
		fmt.Println(key,value)
	}
}

func changevalue(citymap map[string]string){
	citymap["USA"] = "FUCK"
}

func main(){
	citymap := make(map[string]string)
	//添加k-v
	citymap["China"] = "Beijing"
	citymap["Japan"] = "Tokyo"
	citymap["USA"] = "Newyork"

	//遍历
	changevalue(citymap)
    printmap(citymap)
	//删除
	delete(citymap, "China")

	//修改
	citymap["USA"] = "DC"

	for key, value := range citymap{
		fmt.Println(key,value)
	}

	fmt.Println(citymap)


}