package main

import "fmt"

type Siswa struct{ //inisiasi struct
	Id int
	Nama string
	NilaiRata int
	Status string 
}


func Lulus(user Siswa){ //tanpa reference 
	user.Status = "Lulus"
	fmt.Print(
		"Status: ", user.Status, "\n",
	)
}

func LulusFixed(user *Siswa){ //Deference 
	user.Status = "Lulus"
	fmt.Print(
		"Status: ", user.Status, "\n",
	)
}

func main(){
	siswa1 := Siswa{1, "Bagas Arisandi", 99, "Kelas 3"}

	fmt.Print(
		"Id: ", siswa1.Id, "\n",
		"Nama: ", siswa1.Nama, "\n",
		"Rata-Rata: ", siswa1.NilaiRata, "\n",
		"Status: ", siswa1.Status, "\n",
	)

	fmt.Println("\nFunc Lulus:")
	Lulus(siswa1)
	fmt.Println("Dirubah:\n", siswa1.Nama, siswa1.Status)
	Lulus(siswa1)
	fmt.Println("\nFunc LulusFixed:")
	LulusFixed(&siswa1) //reference ke struct Siswa
	fmt.Println("Dirubah:\n",siswa1.Nama, siswa1.Status)


}