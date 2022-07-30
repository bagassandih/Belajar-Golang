package main 

import (
	"fmt"
	"errors"
)

/*
1. Buat func sum untuk menjumlahkan value array/slice tersedia
*/
func sum(angka []int) (hasil int) {
	for _, daftarAngka := range angka {
		hasil += daftarAngka
	}
	return
}



/*
2. Buat func calculate untuk menkalkulasikan value 
dengan operator yang ditentukan beserta handling error
*/

func calculate(angkaSatu int, angkaDua int, operator string) (hasil int, cekError error){
	switch operator{
	case "+": 
		hasil = angkaSatu + angkaDua
	case "-":
		hasil = angkaSatu - angkaDua 
	case "*":
		hasil = angkaSatu * angkaDua 
	case "/":
		hasil = angkaSatu / angkaDua 
	default:
		cekError = errors.New("operator tidak tersedia")
		}
	return hasil, cekError
	fmt.Println(hasil)
	
}

func main(){
// 1. func sum
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)
	fmt.Println()

//2. func calculate
	hasil, err := calculate(10,2,"x")
	if err != nil {
		fmt.Println("Kesalahan")
		fmt.Println(err.Error())
	}
	
	// calculate(10,2,"-")
	// calculate(10,2,"*")
	// calculate(10,2,"/")
}