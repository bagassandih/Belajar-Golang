package main

import "fmt"

func main(){
	
	//kotak
	for i := 0; i < 8; i++{ // kebawah
		for j :=0; j < 8; j++ { // kesamping
		fmt.Print("*")
		}
		fmt.Println()
	}

	//segitiga siku
	for i := 0; i <= 8; i++{ // kebawah
		for j :=0; j < i; j++ { // kesamping
		fmt.Print("*")
		}
		fmt.Println()
	}




}