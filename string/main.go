package main

import (
	"fmt"
	"strings"
)

const (
	str = "something"
	substr = "some"
)

func main() {

// len => panjang atau jumlah string
	kalimat := "Hello"
	lenkalimat := len(kalimat)
	fmt.Println(lenkalimat)
	fmt.Println(len("Test berapa panjang huruf ini"))
	fmt.Println()
	
// compare string atau menyamakan string
	str1 := "xyz"
    str2 := "xys"
    fmt.Println(str1 == str2)	
    fmt.Println("xys" == str2)
    fmt.Println()

 // contains cek apakah terdapat string yg dicari
 // import strings
 // string pakai s
 res:= strings.Contains(str, substr)
 fmt.Println(res)
 fmt.Println(strings.Contains(str1, "xyzzxzs"))
 fmt.Println()    


// substring untuk mengambil string dari index tertentu
kata:= "malamwminggu"
subkata:= kata[5:len(kata)]
fmt.Println(subkata)
fmt.Println( kata[6:len(kata)] )
fmt.Println()  

// replace
ini := "Hari ini cuma dapat satxu"
itu := strings.Replace(ini, "x", "", -1)
fmt.Printf("%s\n", itu)
fmt.Println( strings.Replace(ini, "satxu", "dua belas", -1) )

// insert string
warna := "Merah"
index := len(warna)
warna_replace := warna[:index] + " Muda" + warna[index:]
fmt.Println(warna, warna_replace)
fmt.Println(warna[:index] +" Tua")

// split di bagi ke dalam array
nama := "Bagas Arisandi Hidayat"
fmt.Println( strings.Split(nama, " ") )

// ToLower string menjadi huruf kecil
fmt.Println( strings.ToLower(nama) )

// ToUpper &ToTitle string menjadi huruf besar
fmt.Println( strings.ToUpper(nama) )

// Trim biasa digunakan untuk parsing menghilangkan spasi di kiri dan kanan
fmt.Println( strings.Trim("  Bagas Arisandi     ", " ") )

//ReplaceAll replace semuanya 
fmt.Println( strings.ReplaceAll("Bagis Arisandi\nBagis Hidayat", "Bagis", "Bagas") )


}