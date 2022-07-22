package Kalkulasi

//pakai huruf besar agar bisa di import dari luar(GLOBAL)
func Tambah(AngkaSatu int, AngkaDua int) int {
	return tambah(AngkaSatu, AngkaDua)
}

//pakai huruf kecil untuk lokal(private)
func tambah(AngkaSatu int, AngkaDua int) int {
	return AngkaSatu + AngkaDua
}