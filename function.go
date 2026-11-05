package main
import "fmt"

func main () {
	//If Else Function

	name := "Charles"

	name = "joko"

	if name == "Charles" {
		fmt.Println("Hello ", name)
	} else if name == "joko" {
		fmt.Println("Oh Hello", name)
	} else {
		fmt.Println("?")
	} 

	//If Short Statement 	

	if length := len(name); length < 5 {
		fmt.Println("less than 5")
	} else {
		fmt.Println("namanya sudah oke ")
	}


	// Switch Expression

	switch name {
	case "charles":
		fmt.Println("yo ", name)
	case "joko":
		fmt.Println("yo", name)
	default:
		fmt.Println("?")
	}

	// switch short statement

	switch length := len(name); length < 5 {
	case true:
		fmt.Println("x")
	case false:
		fmt.Println("V")
	}
	
	name = "jawa ngehits"
	length := len(name)
	switch {
	case length < 5:
		fmt.Println("x")

	case length > 5:
		fmt.Println("V")
	
	default:
		fmt.Println('O')
	}



	
}