package main

import "fmt"

func main(){

	//forloop
	for i:=1; i<=10; i++{
		fmt.Println("Testing ForLoop ke - ", i)
	}
	
	//whileloop (bkn sintaks khusus go)
	//biasa untuk start loop tanpa henti jika tanpa kondisi tertentu
	Ulang := 1
	for Ulang <= 10 {
		fmt.Println("Testing WhileLoop ke - ", Ulang)
		Ulang++
	}


	//for each
	//contoh cek huruf tiap urutan(index)

	// for index, letter := range title{
	// 	fmt.Println("index: ", index, " letter: ", string(letter) )
	// }
	
	title := "Golang the best language"
	for _, letter := range title{
		fmt.Println( " letter: ", string(letter) )
	}

	// _ adalah variabel dump(menampung nilai yg tidak digunakan )



}