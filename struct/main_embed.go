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

func tampilKelas(Kelas kelas){ //typenya akses ke struct kelas
	fmt.Printf("Kelas: %s", Kelas.namaKelas)
	fmt.Println()
	fmt.Printf("Ketua: %s", Kelas.ketuaKelas.namaPanggilan)
	fmt.Println()
	fmt.Printf("Jumlah: %d\n", len(Kelas.paraSiswa) )
	for i, daftarParaSiswa := range Kelas.paraSiswa{
		fmt.Println(i+1, daftarParaSiswa.namaDepan + " " +daftarParaSiswa.namaBelakang)
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
	tampilKelas(Kelas)
	// getter dengan function
	
}