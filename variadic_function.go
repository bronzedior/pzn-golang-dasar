package main

import "fmt"

func sumAll(numbers ...int) int {	//non-fixed number of arguments
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}