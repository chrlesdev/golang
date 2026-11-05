package main

import "fmt"

func main () {
	person := map[string]string{
		"name" : "eko",
		"address": "subang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	books := make(map[string]string)
	books["title"] = "belajar Go-Lang"
	books["author"] = "Charles"
	books["Released Year"] = "2026"

	// delete(books, "Released Year")

	fmt.Println(books)
}