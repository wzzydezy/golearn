package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {

		defer fmt.Println("sub end")

		fmt.Println("running")

		c <- 666
	}()

	num := <-c

	fmt.Println(num)
	fmt.Println("main end")

}
