package main 

import (
	"fmt"
	"errors"
	//import dulu
)

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
	return
	
}

func main(){

//2. func calculate
	hasil, err := calculate(10,2,"x")
	if err != nil {
		fmt.Println("Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(hasil)
	
	// calculate(10,2,"-")
	// calculate(10,2,"*")
	// calculate(10,2,"/")
}