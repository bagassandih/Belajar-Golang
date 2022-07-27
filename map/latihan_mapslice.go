package main 

import "fmt"

func main(){

//	data siswa beserta nilai
	siswa := []map[string]string {
		{"nama": "Abdul Fatah", "grade": "A" },
		{"nama": "Alfian", "grade": "B" },
		{"nama": "Amalia Ihtiarini", "grade": "C" },
		{"nama": "Bagas Arisandi", "grade": "A" },
		{"nama": "Bayu Ade", "grade": "C" },
		{"nama": "Cahya Aji", "grade": "B" },
	}
// slice tampung untuk nama siswa yang remedial
	siswaRemed := []string {}

// tampilkan data siswa
	fmt.Println("==========Daftar Siswa==========")
	for _, daftarSiswa := range siswa {
		fmt.Println(daftarSiswa["nama"])
	}

// scanning
	for _, daftarSiswa := range siswa {
		// jika siswa gradenya C
		if daftarSiswa["grade"] == "C"{
		//	masukkan ke slice siswaRemed
			siswaRemed = append(siswaRemed, daftarSiswa["nama"])
			}
		}
// tampilkan yang remedial
	fmt.Println("=====Daftar Siswa Remedial======")
	for _, daftarSiswaRemed := range siswaRemed {
		fmt.Println(daftarSiswaRemed)
		}
	}