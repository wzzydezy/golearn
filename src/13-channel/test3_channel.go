package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//close(c)
		}

		close(c)
	}()

	// for {
	// 	if data, ok := <-c; ok {
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main end")
}
