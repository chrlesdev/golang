package main

import "fmt"

func main () {
	var name string
	name = "ko edi"
	fmt.Println(name)
	name = "eko cuy"
	fmt.Println(name)
	const age = 20
	fmt.Println(age)

	var (
		firstName = "charles"
		lastName = "cornelio"
	)

	var nilai32 int32 = 32768
	var nilai64	int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) 
	

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(firstName,lastName)
	// age = 30 
	// fmt.Println(age)
	
	// fmt.Println("hello, world!")
	// fmt.Println(len("aasldaddsas"))
	// fmt.Println("abcdefghijkl"[5])
}