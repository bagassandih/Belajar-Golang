package main
import "fmt"

/*
	func <nama>(){statement}
	func main() {}

	func <nama>()<tipe_return> {statemen}
	func phi()float64 { return 3.14 }

	func <nama>(<parameter> <tipe>)<tipe_return> {statemen}
	func persegi(angka int) int{return angka*angka}
*/

// tanpa parameter
func ngomong() {
	fmt.Println("Hello")
}

// dengan parameter
func sholat(angka int){
	var status string

	switch angka{
	case 5:
		status = "Subuh"
	case 12:
		status = "Zuhur"
	case 15:
		status = "Ashar"
	case 18:
		status = "Maghrib"
	default:
		status = "Isya"
	}

	fmt.Println("Waktunya Sholat", status)
}


// return 
	func luasPersegi(sisi int) int {
		return sisi * sisi
	}

// multiple return
	func persegiPanjang(panjang int, lebar int) (int, int) {
		var luas, keliling int
		
		luas = panjang * lebar
		keliling = 2*(panjang + lebar)
		fmt.Println("Luas Persegi Panjang:", luas)
		fmt.Println("Keliling Persegi Panjang:", keliling)
		return luas, keliling
	}	

//	return dengan nama
	func perkalian(angkaSatu int, angkaDua int) (hasil int){
		hasil = angkaSatu * angkaDua
		return
	}

func main() {
	
	ngomong()
	
	sholat(12)

	fmt.Println("Luas Persegi:", luasPersegi(3) )

	persegiPanjang(3,5)

	fmt.Println("Perkalian:", perkalian(5,10) ) 

}