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

// replace
ini := ""

}