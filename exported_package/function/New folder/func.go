package function
import "fmt"

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

func (murid siswa) tampil(){
	fmt.Printf("Nama: %s %s \nNickname: %s \nNilai: %d \nHadir: %t\n", murid.namaDepan, murid.namaBelakang, murid.namaPanggilan, murid.nilai, murid.ada)
} 
