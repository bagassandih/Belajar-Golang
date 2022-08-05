package main

import "fmt"

func GetFactorial(n int) int {
	rekursif := n - 1
	if n == 0 { return 1 
	} else {
		return n * GetFactorial(rekursif)
	}
}

func GetFibo(n int) []int {
	if n == 2 { return []int{1, 1} }
	Fibo := GetFibo(n-1)
	FiboNext := Fibo[len(Fibo)-1] + Fibo[len(Fibo)-2]
	return append(Fibo, FiboNext)
}


func main(){

	
	fmt.Println(GetFactorial(5))
	fmt.Println(GetFibo(5))

}