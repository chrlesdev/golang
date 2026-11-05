package main 

import "fmt"

func sayHello(firstName string, lastname string){
	fmt.Println("Hello", firstName, lastname)
}

func getHello(name string) string {
	return "hello " + name 
}

func getFullName() (string, string) {
	return "joko", "jawa"
}

func main () {
	result := getHello("jawa")
	sayHello("joko", "jawa")
	fmt.Println(result)
}