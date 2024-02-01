package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {	//pass by reference
	address.Country = "Indonesia"
}

func main() {
	// var address *Address = &Address{}	
	// ChangeCountryToIndonesia(address)	

	var address Address = Address{}	
	ChangeCountryToIndonesia(&address)	

	fmt.Println(address)
}