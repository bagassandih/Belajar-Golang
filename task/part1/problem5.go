package main
import "fmt"

func main(){
	var T float64 = 20
	var r float64 = 4
	const Pi = 3.14

	fmt.Println( 2*Pi*r*(r + T) )
}