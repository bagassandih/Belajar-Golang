package main

import "fmt"
func main() {

	var prima [5]int = [5]int {2, 3, 5}

	//loop 1	
	fmt.Println("Loop teknik 1")
	for index := 0; index < len(prima); index++ {
		fmt.Println(prima[index])
	}

	//loop2
	fmt.Println("Loop teknik 2")
	for index, daftarPrima := range prima {
		fmt.Println("index", index, " = Valuenya", daftarPrima)
	}
	// dump _ variabel untk menampung tanpa menampilkan
	fmt.Println("Loop teknik 2.5 dengan variabel _")
	for _, value := range prima {
		fmt.Println(value)
	}

	//loop3
	fmt.Println("Loop teknik 3")
	indeks := 0
	for range prima {
		fmt.Println(prima[indeks])
		indeks++
	}


}