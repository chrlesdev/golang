package main 

import "fmt"

func sayHello(firstName string, lastname string){
	fmt.Println("Hello", firstName, lastname)
}

func getHello(name string) string {
	return "hello " + name 
}

func main () {
	result := getHello("jawa")
	sayHello("joko", "jawa")
	fmt.Println(result)
}