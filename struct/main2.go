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
	// inisiasi sekaligus set properti dengan isinya
	siswa1 := siswa{
		id: 1,
		namaDepan: "Andy",
		namaBelakang: "Sansono",
		namaPanggilan: "Andy",
		nilai:	100,
		ada: true,
	} 

	// inisiasi tidak urut atau acak sekaligus set properti dengan isinya
	siswa2 := siswa{
		nilai:	90,
		id: 2,
		namaBelakang: "Bromo",
		namaPanggilan: "Wijaya",
		ada: true,
		namaDepan: "Bro",
	}


	// getter
	fmt.Println(siswa1) // tampilkan semua 

	
	// getter
	fmt.Println(siswa2) // tampilkan hanya properti yang dimaksud

}