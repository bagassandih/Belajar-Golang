package main

import "fmt"

func main(){
	/*
	hitung rata-ratanya dalam bentuk koma kalau memang bentuknya koma
	*/
	nilai := [8]int {100, 80, 75, 92, 70, 93, 88, 67}
	total := 0
 
 	//hitung total dulu
 	for _, daftarNilai := range nilai {
 		total += daftarNilai
 	}

 	//hitung rata rata dan ubah ke float
 	rataRata :=  float64(total) / float64(len(nilai)) 
	fmt.Println(rataRata)


}