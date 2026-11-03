package main

import "fmt"

func main() {
	var names [2]string
	
	names[0] = "joko"
	names[1] = "jawa"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{
		90,
		80,
		20,
		59,
	}

	fmt.Println(len(values))

	//slice


	element := [...]string {"fire", "water", "wind", "earth", "light", "darkness", "ice"}
	
	slice1 := element[2:6]
	fmt.Println("slice 1", slice1)

	slice2 := element[:4]
	fmt.Println("slice 2",slice2)

	slice3 := element[4:]
	fmt.Println("slice 3", slice3)


	//append 

	days := [...]string {"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(daySlice2)

	newSlice := make([]string, 2,5)
	newSlice[0] = "jawa"
	newSlice[1] = "joko"
	// newSlice[2] = "waduh"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

}