package main

import (
	"fmt"
	"reflect"
	)

func main() {
	fmt.Println( "Membuat Slice dari Array" )
	// inisiasi array untuk melakukan slice
	// array[low:high] opsional
	var binatang = [...]string {
		"Kucing",
		"Anjing",
		"Kelinci",
		"Kadal",	
		"Ular",
	}
	// inisiasi slice dari array 
	var binatangReptil []string = binatang[3:5]

	//cek
	fmt.Println( "Reflect var binatang:", reflect.ValueOf(binatang).Kind() )
	fmt.Println( "Reflect var binatangReptil:", reflect.ValueOf(binatangReptil).Kind() )
	fmt.Println( binatangReptil )

	// append untuk menambahkan kedalam array atau slice
	binatangReptil = append(binatangReptil, "Buaya")
	fmt.Println( binatangReptil )

	
	var warna = []string {"Merah", "Kuning", "Hijau"}
	warna = append(warna, "Biru", "Hitam")

	// copy 
	warnaBaru := make([]string, 10) //make dulu
	copy( warnaBaru, warna)
	fmt.Println( warnaBaru )


	//zero slice is nil
	var obat []string
	fmt.Printf("s = %v, len = %d, cap = %d\n", obat, len(obat), cap(obat) )

	if obat == nil { //like a null
		fmt.Println("s is nil, ga ada obat")
	}


	

}