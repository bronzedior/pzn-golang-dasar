package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke ", counter)
	//	counter++
	//}

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("Selesai")

	names := []string{"Eko", "Kurniawan", "Khannedy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println("Nama doang =", name)
	}
}