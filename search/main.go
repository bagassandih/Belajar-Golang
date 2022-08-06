package main

import (
	"fmt"
	"sort"
)


func linier(n []int, x int) int{
	for i := 0; i < len(n); i++{
		if n[i] == x {
			return i
		} 
	}	
	return 0
}




func main(){

var angka = []int{1,3,4,1,2}
var cari = 4
fmt.Println( "Daftar Angka:", angka, "Angka:", cari, "Ada di index:", linier(angka, cari) )
	

// dengan package sort	
elements := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
value := 31
findIndex := sort.SearchInts(elements, value)
fmt.Println(elements)
if value == elements[findIndex] {
	fmt.Println("Value found in elements", findIndex)
} else {
	fmt.Println("Value not found in elements")
}

}