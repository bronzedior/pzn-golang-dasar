package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	donny := Man{"Donny"}
	donny.Married()

	fmt.Println(donny.Name)
}