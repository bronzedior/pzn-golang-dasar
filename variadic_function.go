package main

import "fmt"

func PrependItems(slice []int, values ...int) []int {
	// return append(slice, values...)	//append in the back
	return append(values, slice...)	//append in the front
}

func sumAll(numbers ...int) int {	//non-fixed number of arguments
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	slice := []int{3, 4, 5}
	newSlice := PrependItems(slice, 1, 2, 3)
	fmt.Println(newSlice)

	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}