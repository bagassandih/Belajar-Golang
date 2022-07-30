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





func main(){
// 1. func sum
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)
	fmt.Println()


}