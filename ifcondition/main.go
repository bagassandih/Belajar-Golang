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
	} else if Nilai >= 90 {
		Grade = "B+"
	} else if Nilai >= 80 {
		Grade = "B"
	} else if Nilai >= 70 {
		Grade = "C+"
	} else if Nilai >= 60 {
		Grade = "C"
	} else if Nilai >= 50 {
		Grade = "D"
	} else {
		Grade = "Tidak ada"
	}

	fmt.Println(Grade)

	//bilangan genap/ganjil
	angka := 15
	var status string
	if angka%2==0{
		status = "Genap"
	} else {
		status = "Ganjil"
	} 

	fmt.Printf("Angka %d ini %s \n", angka, status)

	//jam
	jam := 12
	var statusJam string 
	if jam >= 5 {
		statusJam = "Pagi!"
	} else if jam >= 10 {
		statusJam = "Menjelang Siang!"
	} else if jam >= 12 {
		statusJam = "Siang"
	} else if jam >= 15 {
		statusJam = "Menjelang Sore!"
	} else if jam >= 16 {
		statusJam = "Sore!"
	} else if jam >= 19 {
		statusJam = "Menjelang Malam!"
	} else if jam >= 20 {
		statusJam = "Malam!"
	} else if jam >= 23 {
		statusJam = "Menjelang Tengah Malam!"
	} else {
		statusJam = "Menjelang Pagi!"
	}

	fmt.Printf("Selamat %s", statusJam)
}