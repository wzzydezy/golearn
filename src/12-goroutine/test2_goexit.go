package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		//go一个形参为空，返回值为空的一个匿名函数
		go func() {
			defer fmt.Println("A.defer")

			func() {
				defer fmt.Println("B.defer")
				runtime.Goexit()
				fmt.Println("B")
			}()
		}()
	*/

	go func(a int, b int) bool {
		fmt.Println(a, b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
