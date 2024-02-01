package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Fransiscus"
	middleName = "Bronzedior Driandonny"
	lastName = "Noryon"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}