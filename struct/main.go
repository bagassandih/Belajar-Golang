package main
import "fmt"

// inisiasi struct dengan properti default

type siswa struct{
	id 	int
	namaDepan, namaBelakang string
	namaPanggilan string
	nilai int
	ada bool
}

func main(){
	
	siswa1 := siswa{} // inisiasl lagi dengan variabel agar bisa diakses
	
	fmt.Println(siswa1) // tampilkan semua 

	//setter
	siswa1.id = 1
	siswa1.namaDepan = "Aang"
	siswa1.namaBelakang = "Hidayat"
	siswa1.namaPanggilan = "Eng"
	siswa1.nilai = 100
	siswa1.ada = true

	//getter
	fmt.Println(siswa1.namaPanggilan) // tampilkan hanya properti yang dimaksud

}