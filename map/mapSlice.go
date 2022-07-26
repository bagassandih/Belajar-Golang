package main

import (
	"fmt"
	"reflect"
	)

func main() {
// inisiasi slice map
	siswa := []map[string]string{
		//pakai kurung kurawal
		{"nama": "Bagas", "grade": "A"},
		{"nama": "Budi", "grade": "B"},
		{"nama": "Agung", "grade": "C+"},
		{"nama": "Pras", "grade": "A"},
	}

	fmt.Println(reflect.ValueOf(siswa).Kind())
	fmt.Println()

	for _, daftarSiswa := range siswa{
		fmt.Println(daftarSiswa["nama"], " -> ", daftarSiswa["grade"])
	}


}