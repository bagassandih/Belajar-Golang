package main

import "fmt"

func main(){

	
	//variabel
	title := "Golang the best language"
	
	//jika index genap
	fmt.Println("INDEX GENAP")
	for index, letter := range title{
		if index%2==0{
		fmt.Println( "index: ", index, " letter: ", string(letter))
		}
	}


	//vokal konsonan
	fmt.Println("VOKAL KONSONAN CHAR")
	for index, letter := range title{
		if letter == 'a' || letter == 'i' || letter == 'u' || letter == 'e' || letter == 'o' {
		fmt.Println( "index: ", index, " letter: ", string(letter))
		}
	}

	//petik satu ' otomatis menjadi char


	//koversi letter manjadi string
	fmt.Println("VOKAL KONSONAN STRING")
	for index, letter := range title{
		letterString := string(letter)
		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		fmt.Println( "index: ", index, " letter: ", string(letter))
		}
	}

	//dengan switch
	fmt.Println("VOKAL KONSONAN SWITCH")
	for index, letter := range title{
		letterString := string(letter)
		switch letterString{
		case "a","i","u","e","o":
			fmt.Println("index: ", index, " letter: ", string(letter))
			}
	}

	
	


}