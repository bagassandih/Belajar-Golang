package main

import "fmt"

func main(){
	//inisiasi variabel 
	Umur := 24
	Umur = 17

	Nilai := 88
	var Grade string

	//kondisi umur
	if Umur > 18 {
		fmt.Println("Sudah dewasa")
	} else {
		fmt.Println("Belum dewasa")
	}

	//kondisi nilai
	if Nilai == 100 {
		Grade = "A"
	} case Nilai >= 90 {
		Grade = "B+"
	} case Nilai >= 80 {
		Grade = "B"
	} case Nilai >= 70 {
		Grade = "C+"
	} case Nilai >= 60 {
		Grade = "C"
	} case Nilai >= 50 {
		Grade = "D"
	} else {
		Grade = "Tidak ada"
	}

	fmt.Println(Grade)

}