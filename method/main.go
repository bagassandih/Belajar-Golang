package main
import "fmt"

//STRUCT BISA SEBAGAI FIELD (struct didalam struct)
// inisiasi struct 

type siswa struct{
	id 	int
	namaDepan string
	namaBelakang string
	namaPanggilan string
	nilai int
	ada bool
}

type kelas struct{
	namaKelas string
	ketuaKelas siswa //ambil data dari strutc siswa
	paraSiswa []siswa //slice dari struct siswa
	ada bool

}


// Func berubah jadi method
// func (nama nama_struct) nama_method() return {}
 func (classroom kelas) tampil() {
 	fmt.Printf("Kelas: %s \nKetua: %s\nJumlah: %d\n", classroom.namaKelas, classroom.ketuaKelas.namaPanggilan, len(classroom.paraSiswa) ) //panggil properti dari funcnya bukan structnya
 	for i, daftarSiswa := range classroom.paraSiswa{
 		fmt.Println(i+1, daftarSiswa.namaDepan + " " + daftarSiswa.namaBelakang)
 	}
 }



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

	//ambil slice dari para siswa di struct kelas
	paraSiswa := []siswa{siswa1, siswa2, siswa3}
	//inisiasi Kelas dari struct
	Kelas := kelas{"TKJ 2", siswa3, paraSiswa, true}
	

	// getter dengan method
	Kelas.tampil()
	
}