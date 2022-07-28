package main

import (
	"fmt"
	"strconv"
	"strings"
)


func FullPrima (n int) bool {
// convert int ke string
	number := strconv.Itoa(n)
// masukkan ke array dengan split
	strNumber := strings.Split(number, "")
// tampung hasil cek ke array	
	var tampungCek []bool 
// variabel hasil
	var hasilCek bool
// variabel untuk hitung banyak bool, menentukan full prima
	var hitungCek int


// scan index array yang mengandung angka satuan prima
	for _, daftarNumber := range strNumber {
		switch daftarNumber {
		case "2","3","5","7": //daftar angka satuan prima
			tampungCek = append(tampungCek, true) //masuk ke array 
			hitungCek++ //terhitung kena cek
		default:
			tampungCek = append(tampungCek, false) //masuk ke array
		}

	}

// tentukan hasil akhir berdasarkan jumlah true dg index tampung hasil
	if hitungCek == len(tampungCek){
		hasilCek = true
	}

	fmt.Print( //line manual \n untuk tambah line baru
		"Bilangan ", n, " => ", tampungCek, "\n",
		"Jumlah true: ", hitungCek, "\n",
		"FullPrima: ", hasilCek, "\n\n",
	)
	return hasilCek
	//anjay jadi
}

func main(){
	FullPrima(23517)
	FullPrima(2233)
	FullPrima(113)
	FullPrima(57)
	FullPrima(8)
	

}