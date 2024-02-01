package main

import "fmt"

func main() {
	const (
		firstName string = "Fransiscus"
		lastName         = "Noryon"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Budi"
	// lastName = "Joko"
}