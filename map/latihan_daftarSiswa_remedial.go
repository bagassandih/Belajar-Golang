package main

import "fmt"

func main(){

// inisiasi map
	nilaiIPA := map[string]float64{
		"Bagas": 95.5,
		"Andi": 92.3,
		"Budi": 89.7,
		"Chandra": 97.7,
		"Indun": 60.25 ,
		"Alfi": 72.3 ,
		"Linka": 90.5 ,
	}
	
// insiasi KKM atau standar nilai
	kkm := 85.5 
// tampilkan yang melebihi KKM
	for i, daftarNilai := range nilaiIPA {
		if daftarNilai >= kkm {
			fmt.Println(i, "dengan nilai:", daftarNilai, "[ TIDAK REMEDI ]")
			} else {
				fmt.Println (i, "dengan nilai:", daftarNilai, "[REMEDI]")
			}		
		} 


}