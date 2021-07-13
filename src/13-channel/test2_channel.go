package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	fmt.Println(len(c), cap(c))

	go func() {
		defer fmt.Println("sub end")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println(i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num)
	}

	fmt.Println("main end")
}
