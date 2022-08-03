package main
import "fmt"

//insiasi interface
type Luas interface{
	HitungLuas() int //biasanya isinya method
}

// inisiasi struct bangun datar persegi
type Persegi struct{ //implementasi interface luas
	Sisi int
}

// inisiasi method untuk persegi
func (persegi Persegi) HitungLuas() int{
	return persegi.Sisi * persegi.Sisi
}

// inisiasi struct bangun datar persegi panjang
type PersegiPanjang struct {
	Panjang int
	Lebar int
}

// inisiasi method untuk persegi panjang
func (pp PersegiPanjang) HitungLuas() int{
	return pp.Panjang * pp.Lebar
}
 
// inisiasi cek luas 
// pokok permasalahan interface ada dsini agar parameter bisa flexibel
func CekLuas(bangunDatar Persegi) int{ 
// menentukan bangun datarnya masih manual, seharusnya dinamis bisa menyesuaikan walaupun dengan satu function
	return bangunDatar.HitungLuas() //menghubungkan ke interface
}
func main(){
	bangunDatar := Persegi{Sisi: 4}
	luas := CekLuas(bangunDatar) // hitung dengan interface sesuai bangun datarnya, yaitu persegi
	fmt.Println("Luas:", luas)

}