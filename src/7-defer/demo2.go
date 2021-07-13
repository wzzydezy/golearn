package main
import "fmt"

// return 比 defer 早执行
func deferfunc(){
	fmt.Println("D")
}

func returnfunc() int {
	fmt.Println("R")
	return 0
}

func returnandefer() int {
	defer deferfunc()
	return returnfunc()
}

func main(){
	returnandefer()
}