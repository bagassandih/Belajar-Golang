package main

import "fmt"

func main(){
	//inisiasi variabel 
	Nomor := 5

	//case digunakan untuk kondisi spt pilihan yg banyak
	//dengan expresion
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
	//tanpa expression
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

	//fallthorugh, jika kondisi true akan menampilkan kondisi dibawahnya juga
	number := 1
	var statusNumber, tambahan string
	switch number {
	case 1:
		statusNumber = "Number adalah satu"
		tambahan 	 = "Ganjil"	
		fallthrough //setelah case 1 di eksekusi akan lanjut ke kondisi dibawahnya
	case 2:
		statusNumber = "Number adalah dua"
		tambahan 	 = "Genap"	
		fallthrough //setelah case 2 di eksekusi akan lanjut ke kondisi dibawahnya
	case 3:
		statusNumber = "Number adalah tiga"
		tambahan 	 = "Ganjil"	
	default:
		statusNumber = "Number adalah kosong"
		tambahan	 = "Tidak ada"
	}
	fmt.Println(statusNumber, tambahan)
}