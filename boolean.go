package main	

import "fmt"

func main() {
	var nilaiAkhir int= 90
	var absensi int = 80

	// var lulusNilaiAkhir bool = nilaiAkhir > 80
	// var lulusAbsensi bool = absensi > 80

	// var lulus bool = lulusAbsensi 

	if (nilaiAkhir && absensi > 80) {
		fmt.Println("selamat kamu Lulus")
	} else {
		fmt.Println("dawg, man NT yak")
	}

	// fmt.Println(lulus)
}