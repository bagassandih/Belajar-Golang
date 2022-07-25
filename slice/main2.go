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


// long deklarasi
var even_number []int
fmt.Printf( "elements = %v, len = %d, cap = %d \n", even_number, len(even_number), cap(even_number) )

// long deklarasi dengan valuenya
var odd_number =  []int{1,3,4,5,6,8}
fmt.Printf( "elements = %v, len = %d, cap = %d \n", odd_number, len(odd_number), cap(odd_number) )

// short deklarasi dengan value
number:= []int{10,11,12,13,14}
fmt.Println("elements =", number, "len =", len(number), "cap =", cap(number) )

//dengan fungsi make
var prima = make([]int, 5, 10) //buat 5 len dengan 10 cap
	fmt.Printf("elements = %v, len = %d, cap = %d\n", prima, len(prima), cap(prima) )
	


	
	

}