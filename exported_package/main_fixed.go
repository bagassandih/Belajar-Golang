package main
import (
	"fmt"
	"main/function"
	)

/*
File ini file fixed, berhasil import module nya dan terhubung ke main
Perintah atau sintaks  dari package lain sudah berkapital 
setiap perintah yang ada di package lain harus disebutkan didepan sintaksnya 
contoh function.Siswa
*/


func main(){
	
	siswa1 := function.Siswa{
		1,
		"Andy",
		"Sansono",
		"Andy",
		100,
		true,
	} 

	siswa2 := function.Siswa{ 2, "Bromo", "Wijaya", "Bro", 98, true } 
	siswa3 := function.Siswa{ 3, "Candra", "Oktavian", "Candra", 94, true } 

	
	paraSiswa := []function.Siswa{siswa1, siswa2, siswa3}
	Kelas := function.Kelas{"TKJ 2", siswa3, paraSiswa, true}
	

	// getter dengan method
	fmt.Println("======================")
	Kelas.Tampil()
	fmt.Println("======================")
	siswa1.Tampil()

}