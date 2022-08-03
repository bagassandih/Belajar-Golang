package main

import "fmt"

type Remedial interface{
	CekRemedial()
}

type Siswa struct{
	Id int
	Nama string
	Kelas string
	RataNilai float64
	Remedial bool
}

type Kelas struct{
	Id int
	Nama string
	DaftarSiswa []Siswa
}

type Guru struct{
	Id int
	Nama string
	Mapel string
	RataNilai float64
	Remedial bool
}

type Kantor struct{
	Id int
	Nama string
	DaftarGuru []Guru
}

func (class Kelas)CekRemedial(kkm float64){


	// Cara 1
	// var index = 0
	// for i, siswaRemed := range class.DaftarSiswa{
	// 	if siswaRemed.RataNilai < kkm {
	// 		index = i
	// 	}
	// }
	// class.DaftarSiswa[index].Remedial = true


	// Cara 2
	for i, SiswaRemedial := range class.DaftarSiswa{
		if SiswaRemedial.RataNilai < kkm {
			//SiswaRemedial.Remedial = true // not work tidak bisa di index
			class.DaftarSiswa[i].Remedial = true // harus ada indexnya 
		}
	}

	
}


func main(){
    var KKM_Siswa = 82.50
	Bagas := Siswa{2, "Bagas Arisandi", "TKJ 2", 95.25, false}
	Candra := Siswa{3, "Candra Gunawan", "TKJ 2", 92.33, false}
	Amila := Siswa{1, "Amila Anoda", "TKJ 2", 80.25, false}
	Doni := Siswa{4, "Doni Irawan", "TKJ 2", 91.25, false}
	siswaTKJ2 := []Siswa{Bagas, Candra, Amila, Doni}
	TKJ2 := Kelas{2, "TKJ 2", siswaTKJ2}

	// var KKM_Guru = 80.00
	// Mursid := Guru{1, "Mursid Asnawi" , "B. Inggris", 88, false}
	// Ando := Guru{2, "Ando Sudarso", "TIK", 76.66, false}
	// Yuli := Guru{3, "Yuli Yuliana", "Seni Budaya", 99.99, false}
	// GuruSMK := []Guru{Mursid, Ando, Yuli}
	// GuruMapel := Guru{1, "Kantor", GuruSMK}

	TKJ2.CekRemedial(KKM_Siswa)
	// GuruMapel.CekRemedial(KKM_Guru)
	
	

	for _, siswa := range TKJ2.DaftarSiswa {
		fmt.Println(siswa.Nama, siswa.RataNilai, siswa.Remedial)
	}

	// for _, guru := range GuruMapel.DaftarGuru {
	// 	fmt.Println(GuruMapel.Nama, GuruMapel.RataNilai, GuruMapel.Remedial)
	// }
	
	


	

}