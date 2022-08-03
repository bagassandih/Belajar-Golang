package main
import "fmt"

//insiasi interface
type Luas interface{
	HitungLuas() int //biasanya isinya method 
}

// inisiasi struct bangun datar persegi
type Persegi struct{ 
	Sisi int
}

// inisiasi method untuk persegi yang terhubung ke interface dengan parameter hitung luas 
func (persegi Persegi) HitungLuas() int{
	return persegi.Sisi * persegi.Sisi
}

// inisiasi struct bangun datar persegi panjang
type PersegiPanjang struct {
	Panjang int
	Lebar int
}

// inisiasi method untuk persegi panjang yang juga terhubung ke interface dengan parameter yang sama
func (pp PersegiPanjang) HitungLuas() int{
	return pp.Panjang * pp.Lebar
}
 
// inisiasi cek luas 
// pokok permasalahan terpecahkan, sudah flexibel
func CekLuas(bangunDatar Luas) int{ // ganti tipenya dengan nama interfacenya
	return bangunDatar.HitungLuas() // akes apapun bangun datar yang sudah di inisiasi structnya kemudian akses interface yang isinya hitungluas
}
func main(){
	persegi := Persegi{Sisi: 4}
	luasPersegi := CekLuas(persegi) // hitung dengan interface sesuai bangun datarnya, yaitu persegi
	fmt.Println("Luas Persegi:", luasPersegi)

	persegiPanjang := PersegiPanjang{Panjang: 4, Lebar: 2}
	luaspersegiPanjang := CekLuas(persegiPanjang) // hitung dengan interface sesuai bangun datarnya, yaitu persegi panjang
	fmt.Println("Luas Persegi Panjang:", luaspersegiPanjang)
}

/*
Kesimpulan, Interface digunakan untuk mempersingkat method atau function agar lebih mudah dan flexibel
dalam mengaksesnya di func main. Tanpa perlu membuat func yang banyak, walau berbeda tipe datanya.
contoh seperti diatas, setiap bangun datar(struct) yang sudah di inisiasi (persegi dan persegi panjang) 
akan terhubung ke method masing-masing, dimana method juga terhubung ke interface.
*/