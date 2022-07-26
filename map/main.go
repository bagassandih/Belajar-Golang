package main

import (
	"fmt"
	"reflect"
	)

func main() {
	
//	map[tipedata_key]tipedata_value{array}

//	long deklarasi
	var contoh = map[string]int{}
	fmt.Println(contoh)
	fmt.Println(reflect.ValueOf(contoh).Kind())

	var nilaiTik = map[string]int{"Bagas": 100, "Andi": 90}
	fmt.Println(nilaiTik)

//	short deklarasi
	nilaiMtk := map[string]int{}
	fmt.Println(nilaiMtk)

// pakai perintah make
	var nilaiBIndo = make( map[string]int )
	fmt.Println(nilaiBIndo)

// akses map

//	menambahkan / asignment
//	variabel[key] = value
nilaiBIndo["Bagas"] = 100
nilaiBIndo["Joko"] = 99
nilaiBIndo["Bona"] = 80

//	hapus dengan key
delete(nilaiBIndo, "Joko")
fmt.Println(nilaiBIndo)

//	cek ada atau tidak
key, cek:= nilaiBIndo["Bona"]
fmt.Println(key, cek)



}