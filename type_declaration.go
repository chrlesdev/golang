package main

import "fmt"

func main() {
	type noKTP string
	var ktp noKTP = "1231231231313"

	var contoh string = "22222222222"
	var contohKtp noKTP = noKTP(contoh)


	fmt.Println(ktp)
	fmt.Println(contohKtp)
}