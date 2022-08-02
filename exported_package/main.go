package main
import (
	"fmt"
	"main/function"
	)

/*
File ini jika di run terjadi error
karena perintah dan berbagai sintaks yang ada di package lain harus menggunakan
huruf Kapital didepan dan juga harus dilakukan pemanggilan packagenya di setiap perintah
yang berada diluar package 
Silahkan ke file fixed untuk belajar lebih lanjut
*/


func main(){
	
	siswa1 := siswa{
		1,
		"Andy",
		"Sansono",
		"Andy",
		100,
		true,
	} 

	siswa2 := siswa{ 2, "Bromo", "Wijaya", "Bro", 98, true } 
	siswa3 := siswa{ 3, "Candra", "Oktavian", "Candra", 94, true } 

	paraSiswa := []siswa{siswa1, siswa2, siswa3}
	Kelas := kelas{"TKJ 2", siswa3, paraSiswa, true}
	

	// getter dengan method
	fmt.Println("======================")
	Kelas.tampil()
	fmt.Println("======================")
	siswa1.tampil()

}