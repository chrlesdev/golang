package main
import "fmt"

func main () {
	//For Loops

	// counter :=  1

	// for counter <= 5 {
	// 	fmt.Println("counter",counter)
	// 	counter++
	// }

	// for minus := 5; minus >= 1; minus-- {
	// 	fmt.Println("minus",minus)
	// }

	// names := []string {"jawa","jawa","jawa"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// change index to _ if index is not needed

	// for index, name := range names {
	// 	fmt.Println(index,"=" ,name)
	// }


	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("counter", [i] )
	// }


	//break


	for i := 0;  i <= 10; i++ {
		if(i == 5) {
			fmt.Println(i)
			break;
		}
	}
}