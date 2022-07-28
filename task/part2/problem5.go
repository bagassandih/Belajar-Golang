package main
import "fmt"


func pangkat(base, pangkat int) int {
	var hasil = 1 //defaultnya 1 supaya kalo dikali tidak jadi 0
	for i := 0; i < pangkat; i++{
		hasil = hasil * base
	}
	fmt.Println(hasil)
	return hasil
}


func main(){
	pangkat(2 ,3)
	pangkat(5 ,3)
	pangkat(10 ,2)
}