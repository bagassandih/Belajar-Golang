package main
import (
	"fmt"
	"strconv"
	//"strings"
)


func FullPrima(n int) bool {
	var hasil bool
	var convertN = strconv.Itoa(n)
	var tampungHasil []bool
	var tampungCek int
	for i := 0; i < len(convertN); i++ {
		switch string(convertN[i]){
			case "2", "3", "5", "7":
				tampungHasil = append(tampungHasil, true)
				tampungCek++
			default:
				tampungHasil = append(tampungHasil, false)
		}
	}

	if tampungCek == len(tampungHasil){
		hasil = true
	}

	fmt.Println(n, hasil )
	return hasil
}


func main(){
	FullPrima(2)
	FullPrima(3)
	FullPrima(11)
	FullPrima(13)
	FullPrima(23)
	FullPrima(2222) //bisa sampe ribuan

	
}