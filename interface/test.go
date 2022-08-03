package main

import "fmt"

type Siswa struct{
	Nama string
}

func main(){
	siswa := Siswa{"Bagas"}
	fmt.Println(&siswa.Nama)

}