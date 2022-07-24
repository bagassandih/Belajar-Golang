package main

import (
	"fmt"
	"strings"
	)

// //sensor kata berdasarkan index string	
// func sensorKataIndex(kata string) string{
// 	var hasil string
// 	for i := 0; i < len(kata)/2; i++ {
// 	hasil = strings.Replace(kata, string(kata[i]), "*", -1) 
// 	kata = hasil
// 	}
// 	return kata
// }


// func sensorKataVokal(kata string) string{
// 	var hasil, vokalCek string
// 	vokal := [5]string {"a", "i", "u", "e", "o"}

	
// 	for _, daftarVokal := range vokal{
// 	cek :=strings.Contains(kata, daftarVokal)
// 	if cek == true {
// 		vokalCek = daftarVokal
// 		hasil = strings.Replace(kata, vokalCek, "*", -1) 
// 		}
// 	}			

// 	kata = hasil
// 	return kata
// }


func cekSensor (kata string) string{

	hasil := kata //default untuk antisipasi jika tidak disensor
	var vokalCek string
	vokal := [5]string {"a", "i", "u", "e", "o"}
	sensor := false
	KataSensor := [...]string { //daftar kata umpatan
		"anjing",
		"asu",
		"bego",
		"goblok",
		"bangsat",
	}
	

// to lower agar menjadi kecil semua dan mudah di scan
	kata = strings.ToLower(kata)	
// cek kata harus disensor atau tidak
	for _, daftarKataSensor := range KataSensor {
		cekKata := strings.Contains(kata, daftarKataSensor)
		if cekKata == true { sensor=true }
	}

// menyensor kata jika ada kata umpatan
	if sensor == true {
	for _, daftarVokal := range vokal{
	cek :=strings.Contains(kata, daftarVokal)
	if cek == true {
		vokalCek = daftarVokal
		hasil = strings.Replace(kata, vokalCek, "*", -1) 
			}

		}
	}	
	// kembalikan huruf besar di awal kata
	Kapital := strings.ToUpper( string(hasil[0]) ) 
	hasil = Kapital + hasil[1:len(hasil)]
	kata = hasil
	return kata
}

func main(){
	// SEKEDAR CONTOH BUKAN UNTUK DI TIRU KATA KATANYA
	// fmt.Println("Sensor Index")
	// fmt.Println( sensorKataIndex("Bangsat") )
	// fmt.Println( sensorKataIndex("Anjing") )
	// fmt.Println( sensorKataIndex("Asu lu") )
	// fmt.Println()

	// fmt.Println("Sensor vokal")
	// fmt.Println( sensorKataVokal("Bangsat") )
	// fmt.Println( sensorKataVokal("Anjing") )
	// fmt.Println( sensorKataVokal("Asu lu") )
	// fmt.Println()

	fmt.Println("Cek Kata Umpatan dengan Sensor Huruf Vokal")
	fmt.Println( cekSensor("Bangsat") )
	fmt.Println( cekSensor("Anjing") )
	fmt.Println( cekSensor("Asu lu") )
	fmt.Println( cekSensor("Serigala berbulu domba") )
	fmt.Println( cekSensor("Malam Minggu ") )
	fmt.Println( cekSensor("Kamu memang bego!") )
	fmt.Println()	

}