package main

import "fmt"

func main(){

	//isi array manual
	var bahasa [5]string
		bahasa[0] = "Go"
		bahasa[1] = "Ruby"
		bahasa[2] = "JS"
		bahasa[3] = "C#"
		bahasa[4] = "Python"
	fmt.Println(bahasa)

	//isi array langsung
	bahasaDua := [5]string {"C++", "PHP", "CodeIgniter", "CSS", "HTML"}
	fmt.Println(bahasaDua)

	//isi array langsung vertical(kebawah)
	bahasaTiga := [5]string {
		"Perl", 
		"Lavarel", 
		"Bootstrap", 
		"CMS", 
		"HTML5", //khusus vertical bagian akhir element harus dikoma
	}
	fmt.Println(bahasaTiga)

	//isi array langsung tanpa jumlah element
	hewan := [...]string {
		"Kucing", 
		"Anjing", 
		"Serigala", 
		"Beruang", 
		"Koala",
		"Jerapah",
		"Semut", //khusus vertical bagian akhir element harus dikoma
	}

	//cek berapa element array
	fmt.Println(hewan)
	cek := len(hewan)
	fmt.Println(cek)


	//looping array
	for index, daftarHewan := range hewan{
		fmt.Println("index: ", index, " hewan ", daftarHewan)

	}

}