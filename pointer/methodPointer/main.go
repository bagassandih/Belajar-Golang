package main

import "fmt"

type Siswa struct{ //inisiasi struct
	Id int
	Nama string
	NilaiRata int
	Status string 
}


func (user *Siswa)Lulus(){ //Deference 
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

	fmt.Println("\nMethod Lulus:")
	siswa1.Lulus() //tanpa perlu reference manual dengan &, karena tipe ini memang refernce ke struct 
	fmt.Println("Dirubah:\n",siswa1.Nama, siswa1.Status)


}