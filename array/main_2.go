package main

import (
	"fmt"
	"reflect"
	)
func main() {
	var countries [5]string

	countries[0] = "Indonesia"
	countries[1] = "Japan"

	fmt.Println(countries[0], countries[1])

	odd_number := [5]int {1, 2, 4, 42, 14} //inisiasi dengan value/element
	var even_number [5]int= [5]int {13, 23, 44} // partial assignment

	// index yang tdk terisi element/value akan berisi default
	fmt.Println( odd_number )
	fmt.Println( even_number )

	prima := [...]int {2, 3, 3}

	fmt.Println( len(prima) ) //hitung banyak element array
	fmt.Println( reflect.ValueOf(prima).Kind() ) // -

	// inisiasi index yg spesifik
	var even_number2 [5]int = [5]int {1,2,3,4: 5}
	//										 ^ index ke 4 di skip
	fmt.Println( even_number2 )

	


}