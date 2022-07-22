package main
import "fmt"

func main(){
	var number int = 120
	//medapatkan tipe data
	fmt.Printf("Tipe Data: %T \n", number)
	fmt.Printf("Value: %v \n", number)

	//disambung
	fmt.Printf("Tipe Data: %T dan Value: %v \n\n", number, number)


//	sprintf digunakan untuk menampung format kedalam variabel 
//	Printf digunakan untuk langsung menampilkan format di console
	var myString string = fmt.Sprintf("Tipe Data: %T dan Value: %v \n", number, number)
	fmt.Println(myString)


// true or false
	fmt.Printf("True/False? %t \n\n", 1 > 2)

// menampilkan binary (base 2)
	fmt.Printf("Binari dari bilangan 10: %b \n\n", 10)

// menampilkan octal (base 4)
	fmt.Printf("Octal dari bilangan 10: %o \n\n", 10)

// 	menampilkan desimal (base 10) sering dipakai
	fmt.Printf("Desimal dari bilangan 10: %d \n\n", 10)

// 	menampilkan hexadeismal (base16)
	fmt.Printf("Hexa dari bilangan 10: %x \n\n", 10)

//	menampikan notasi ilmiah
	fmt.Printf("Hasil: %e \n\n", 2.4121)

// 	menampilkan desimal dibulatkan
	fmt.Printf("Dibulatkan: %f \n\n", 3.9999999999)

// 	menampilkan desimal kesuluruhan tanpa dibulatkan
	fmt.Printf("kesuluruhan: %g \n\n", 3.9999999999)

// menampilkan string default
	fmt.Printf("Bahasa: %s \n\n", "Golang")

// menampikan string dengan kutip 2 ""
	fmt.Printf("Bahasa Kutip Dua: %q \n\n", "Golang")

// menampilkan width dan presisi(digit) default
	fmt.Printf("%f \n\n", 100.00)

// panjang kekiri (spasi)
	fmt.Printf("Bahasa: %12s lang \n\n", "Go")	

// panjang kanan (spasi)
	fmt.Printf("Bahasa: %-12s lang \n\n", "Go")	

// edit jumlah presisi(digit)
	fmt.Printf("Bahasa: %.2f \n\n", 6.42242)

// gabung width dan presisi
	fmt.Printf("Bahasa: %12.2f \n\n", 6.42242)

// padding (tambahan didepan value)
	fmt.Printf("Bahasa: %05d %05d %05d \n\n", 1, 100, 1000)		

// coba dengan Sprintf (ditampung ke variabel)
	var kodeNoUrut string = fmt.Sprintf("%05d %05d %05d", 1, 20, 300)
	fmt.Println("Kode Barang [", kodeNoUrut, "]")


/*
	[Fomat Value]
	%T cek tipe data
	%v cek valuenya
	%t cek boolean
	%b cek binary
	%o cek octal
	%d cek desimal
	%x cek hexa
	%e cek notasi ilmiah (koma koma)
	%f cek desimal dibulatkan
	%g cek desimal kesluruhan
	%s cek string default
	%q string dengan ""

	[+ jarak spasi (width) di kiri]
	%(jumlah)(format)
	%12s 
	12 	=> + 12spasi di kiri
	s 	=> diikuti format value string

	[+ jarak spasi (width) di kanan]
	%-(jumlah)(format)
	%-12s 
	12 	=> + 12spasi di kanan
	s 	=> diikuti format value string

	[Tampilkan digit presisi]
	%.(jumlah)(format)
	%.2f
	.2	=> presisi setelah koma sebanyak 2 digit
	f 	=> diikuti format value desimal bulat

	[Gabung + jarak spasi (width) dengan presisi digit]
	%(jumlah spasi).(jumlah presisi)(format)
	%12.2f
	12	=> + 12spasi dikiri
	.2	=> presisi setelah koma sebanyak 2 digit
	f 	=> format value desimal dibulatkan

	[Padding (tambah jarak value didepan value)]
	%(value depan)(jumlah digit)(format)
	%05d
	0 	=> value depan adalah 0
	5	=> jumalh digit
	d 	=> format value desimal
*/


}