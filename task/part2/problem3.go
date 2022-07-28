package main
import "fmt"


func primeNumber(number int) bool {
	var hasil bool
	for i := 2; i < number; i++ {
		if number%i == 0 {
			hasil = false
			break // harus di break
		} else {
			hasil = true
		}
	}

	return hasil
}


func main(){
	fmt.Println(primeNumber(11))
	fmt.Println(primeNumber(20))
}