package main

import "fmt"

func Ubah(isiLama int, isiBaru int ){
	fmt.Println("Alamat diFunction:", &isiLama, "Isi:", isiBaru)
}

func UbahFix(isiLama *int, isiBaru int ){ //di dereference agar terhubung
	fmt.Println("Alamat diFunctionFix:", isiLama, "Isi:", isiBaru)
}

func UbahFixReturn(isiLama *int, isiBaru int ) (hasil int){ // deference kemudian di return variabel di func main
	fmt.Println("Alamat diFunctionReturn:", isiLama, "Isi:", isiBaru)
	*isiLama = isiBaru
	hasil = *isiLama
	return hasil 
}
func main(){
	var NumberA int = 10
	fmt.Println("Alamat:", &NumberA, "Isi:", NumberA)

	Ubah(NumberA, 12) // akan berbeda alamat karena tidak di reference
	fmt.Println("Hasil dirubah: ", NumberA)

	UbahFix(&NumberA, 12) // alamat akan sama karena direference
	fmt.Println("Hasil dirubah: ", NumberA)

	UbahFixReturn(&NumberA, 20) // alamat akan sama karena direference
	fmt.Println("Hasil dirubah: ", NumberA)


}