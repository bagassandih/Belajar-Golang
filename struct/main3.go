package main
import "fmt"

// inisiasi struct dengan properti default

type siswa struct{
	id 	int
	namaDepan string
	namaBelakang string
	namaPanggilan string
	nilai int
	ada bool
}

func main(){
	// inisiasi sekaligus set dengan isinya
	// dengan syarat harus urut dengan data struct
	siswa1 := siswa{
		1,
		"Andy",
		"Sansono",
		"Andy",
		100,
		true,
	} 

	siswa2 := siswa{ 2, "Bromo", "Wijaya", "Bro", 100, true } 

	// getter
	fmt.Println(siswa1) // tampilkan semua 

	
	// getter
	fmt.Println(siswa2) // tampilkan hanya properti yang dimaksud

}