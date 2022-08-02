package function
import "fmt"

/*
file ini file yang sudah fixed, sudah berkapital dan bisa di hubungkan ke main
cek file error di new folder untuk pembelajaran

*/

type Siswa struct{
	Id 	int
	NamaDepan string
	NamaBelakang string
	NamaPanggilan string
	Nilai int
	Ada bool
}

type Kelas struct{
	NamaKelas string
	KetuaKelas Siswa 
	ParaSiswa []Siswa 
	Ada bool

}

 func (classroom Kelas) Tampil() {
 	fmt.Printf("Kelas: %s \nKetua: %s\nJumlah: %d\n", classroom.NamaKelas, classroom.KetuaKelas.NamaPanggilan, len(classroom.ParaSiswa) ) //panggil properti dari funcnya bukan structnya
 	for i, daftarSiswa := range classroom.ParaSiswa{
 		fmt.Println(i+1, daftarSiswa.NamaDepan + " " + daftarSiswa.NamaBelakang)
 	}
 }

func (murid Siswa) Tampil(){
	fmt.Printf("Nama: %s %s \nNickname: %s \nNilai: %d \nHadir: %t\n", murid.NamaDepan, murid.NamaBelakang, murid.NamaPanggilan, murid.Nilai, murid.Ada)
} 
