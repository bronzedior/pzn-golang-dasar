package main

import "fmt"

func main() {
	names := [...]string{
		"fransiscus",
		"bronzedior",
		"driandonny",
		"noryon",
	}

	slice1 := names[0:4]	//low:high-1
	fmt.Println(slice1)
	
	slice2 := names[:3]		//0:high-1
	fmt.Println(slice2)

	var slice3 []string = names[3:]	//3:high
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)


	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")	//capacity full, sehingga membuat array days baru
	daySlice2[0] = "Sabtu Lama"
	// daysBaru := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Libur Baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Donny"
	newSlice[1] = "Donny"
	// newSlice[2] = "Eko" // error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Donny", "Donny", "Donny", "Donny")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)


	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}