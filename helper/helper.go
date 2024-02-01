package helper //nama package berdasarkan nama folder

import "fmt"

var version = "1.0.0"	//tidak bisa diaccess di luar package
var Application = "golang" //bisa diaccess di luar package

func sayGoodBye(name string) string {	//tidak bisa diaaccess di luar package
	return "Good Bye " + name
}

func SayHello(name string) string { //bisa diaccess di luar package
	return "Hello " + name
}

func Contoh() {
	sayGoodBye("Donny")
	fmt.Println(version)
}