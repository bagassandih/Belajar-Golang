package main

import (
	"fmt"
	//"reflect"
	)

func main() {

/*
len() digunakan untuk menghitung jumlah elemen slice yang ada

cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice. Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len, tapi bisa berubah seiring operasi slice yang dilakukan
*/


// short deklarasi dengan value
number:= []int{10,11,12,13,14,15,16,17}
fmt.Println("elements =", number, "len =", len(number), "cap =", cap(number) )

//slice variabel number
var newNumber []int = number[3:5]
fmt.Printf("elements = %v, len = %d, cap = %d\n", newNumber, len(newNumber), cap(newNumber) )


var newNumberAgain []int = number[3:4]
fmt.Printf("elements = %v, len = %d, cap = %d\n", newNumberAgain, len(newNumberAgain), cap(newNumberAgain) )

	
	

}