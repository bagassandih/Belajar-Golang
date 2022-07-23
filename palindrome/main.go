package main

import "fmt"

func palindrome(input string) bool {
	balik := []byte{}
	for i := len(input) - 1; i >= 0; i-- {
		balik = append(balik, input[i])
	}
	return input == string(balik)

}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("kupu-kupu"))
	fmt.Println(palindrome("lion"))
}