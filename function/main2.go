package main

import "fmt"

// predefined return
// insiasi tipedata return langsung di functionya
// bisa di multiple juga returnya
func kalkukasi(angkaSatu int, angkaDua int ) (tambah int, kali int, bagi float64){
	tambah = angkaSatu + angkaDua
	fmt.Println(tambah)

	kali = angkaSatu * angkaDua
	fmt.Println(kali)

	bagi = float64(angkaSatu) / float64(angkaDua)
	fmt.Printf("%.2f", bagi)
	return 
}


func main(){
	kalkukasi(10,3)
}