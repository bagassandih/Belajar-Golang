package main

import (
	"fmt"
	"strings"
	)

//sensor kata berdasarkan index string	
func sensorKataIndex(kata string) string{
	var hasil string
	for i := 0; i < len(kata)/2; i++ {
	hasil = strings.Replace(kata, string(kata[i]), "*", -1) 
	kata = hasil
	}
	return kata
}


func sensorKataVokal(kata string) string{
	var hasil, vokalCek string
	vokal := [5]string {"a", "i", "u", "e", "o"}

	
	for _, daftarVokal := range vokal{
	cek :=strings.Contains(kata, daftarVokal)
	if cek == true {
		vokalCek = daftarVokal
		hasil = strings.Replace(kata, vokalCek, "*", -1) 
		}
	//fmt.Println(kata, i, daftarVokal, cek, vokalCek, hasil )

	 

	}
				

	kata = hasil
	return kata
}

func main(){
	// SEKEDAR CONTOH BUKAN UNTUK DI TIRU
	fmt.Println("Sensor Index")
	fmt.Println( sensorKataIndex("Bangsat") )
	fmt.Println( sensorKataIndex("Anjing") )
	fmt.Println( sensorKataIndex("Asu lu") )
	fmt.Println()

	fmt.Println("Sensor vokal")
	fmt.Println( sensorKataVokal("Bangsat") )
	fmt.Println( sensorKataVokal("Anjing") )
	fmt.Println( sensorKataVokal("Asu lu") )

	

}