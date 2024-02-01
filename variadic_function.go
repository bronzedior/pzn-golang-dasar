package main

import "fmt"

func PrependItems(slice []int, values ...int) []int {
	// return append(slice, values...)	//append in the back
	return append(values, slice...)	//append in the front
}

func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
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

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice:", mySlice)
	mySlice = RemoveItem(mySlice, 2)
	fmt.Println("Modified Slice:", mySlice)

	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}