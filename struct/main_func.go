package main
import "fmt"

//STRUCT BISA SEBAGAI PARAMETER FUNCTION
// inisiasi struct dengan properti default

type siswa struct{
	id 	int
	namaDepan string
	namaBelakang string
	namaPanggilan string
	nilai int
	ada bool
}

func displaySiswa (siswa siswa) string { //type siswa
	return fmt.Sprintf("Nama Lengkap: %s %s, Nickname: %s, Nilai: %d", siswa.namaDepan, siswa.namaBelakang, siswa.namaPanggilan, siswa.nilai)
	// Sprint untuk menampung semuanya menjadi string dalam satu kelompok
}

func displaySiswaX (siswa siswa) { //type siswa
	 fmt.Printf("Nama Lengkap: %s %s, Nickname: %s, Nilai: %d", siswa.namaDepan, siswa.namaBelakang, siswa.namaPanggilan, siswa.nilai)
	 // coba dengan printf tanpa parameter return
}

func displaySiswaZ (siswa siswa){ //type siswa
	 fmt.Println("Nama Lengkap: %s %s, Nickname: %s, Nilai: %d", siswa.namaDepan, siswa.namaBelakang, siswa.namaPanggilan, siswa.nilai)
	 // coba dengan println tanpa parameter return
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

	siswa2 := siswa{ 2, "Bromo", "Wijaya", "Bro", 98, true } 


	
	// getter dengan function
	fmt.Println("Dengan Sprintf:")
	fmt.Print(displaySiswa(siswa1) + "\n", displaySiswa(siswa2) + "\n")
	
	fmt.Println("Dengan Printf:")
	displaySiswaX(siswa1)
	fmt.Println()

	fmt.Println("Dengan Println:")
	displaySiswaX(siswa2)

}