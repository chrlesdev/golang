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

	fmt.Println(firstName,lastName)
	conversions()
}