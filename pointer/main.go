package main

import "fmt"

func main(){
	NumberA := 5
	NumberB := &NumberA // reference, ambil alamat number A 
	var NumberC *int = &NumberA // alternatif reference

	fmt.Println("Number A:", NumberA)
	fmt.Println("Number B:", NumberB) // akan menampilkan alamat
	fmt.Println("Number C:", NumberC)
	fmt.Println("Isi Number B:", *NumberB) // * adalah dereference, akan menampilkan isi dari alamatnya
	fmt.Println("Isi Number C:", *NumberC)

	fmt.Println("Ubah isi Number B menjadi 10")
	*NumberB = 10 // ubah isi/value dengan dereference dulu

	fmt.Println("Number A:", NumberA)
	fmt.Println("Number B:", *NumberB)
	fmt.Println("Number C:", *NumberC)

}