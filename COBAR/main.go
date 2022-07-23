package main

import "fmt"

func pangkat (base int, pangkat int) int {
	var result int = 1
	for i := 1; i <= pangkat; i++ {
		result *= base 

	}
	return result
}

func main(){
	fmt.Println( pangkat(2,3) )
	fmt.Println( pangkat(2,4) )
	fmt.Println( pangkat(2,5) )
	fmt.Println( pangkat(2,6) )
}