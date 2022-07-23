package main

import "fmt"

func main(){
	kalimat := "Bagas"

	for i := 0; i < len (kalimat); i++{
		fmt.Println( kalimat[i], "\t>\t", string(kalimat[i]) ) 
	//				^ hasilnya byte           ^ convert ke string
	} 

	for index, huruf := range kalimat {
		fmt.Println("index: ", index, "adalah byte ", huruf, string(huruf) )
	}


}