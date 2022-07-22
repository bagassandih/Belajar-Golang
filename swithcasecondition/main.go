package main

import "fmt"

func main(){
	//inisiasi variabel 
	Nomor := 5

	//case digunakan untuk kondisi spt pilihan yg banyak

	switch Nomor {
	case 1:
		fmt.Println("Nomor Satu")
	case 2:
		fmt.Println("Nomor Dua")
	case 3:
		fmt.Println("Nomor Tiga")
	default:
		fmt.Println("Tidak tersedia, hanya 1 - 3 ")

	}
	

	//contoh kasus di ifcondition
	Nilai := 99
	var Grade string
	switch {
	case Nilai == 100: 
		Grade = "A"
	case Nilai >= 90: 
		Grade = "B+"
	case Nilai >= 80:
		Grade = "B"
	case Nilai >= 70:
		Grade = "C+"
	case Nilai >= 60:
		Grade = "C"
	case Nilai >= 50:
		Grade = "D"
	default: 
		Grade = "Tidak ada"
	}

	fmt.Println(Grade)
}