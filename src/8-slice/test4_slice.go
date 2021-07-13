package main

import "fmt"

func main(){
	var numbers = make([]int, 3, 5)

	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)

	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)

	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3)

	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)
}
