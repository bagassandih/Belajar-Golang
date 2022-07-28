package main
import "fmt"


func palindrome(str string) bool {
	var balik string
	var hasil bool

	for i := len(str)-1; i >= 0; i-- {
		balik = balik + string(str[i])
	}

	if str == balik {
		hasil = true
	}

	fmt.Println(str + " > " + balik + " = ", hasil)
	return hasil

}


func main(){
	palindrome("civic")
	palindrome("kasur rusak")
	palindrome("kupu-kupu")
	palindrome("apa")
	palindrome("lion")
}