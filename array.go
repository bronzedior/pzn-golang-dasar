package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Fransiscus"
	names[1] = "Bronzedior"
	names[2] = "Driandonny"
	names[3] = "Noryon"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
	}

	var values2 = [10]string{
		"Fransiscus",
		"Bronzedior",
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	fmt.Print(values2)
}