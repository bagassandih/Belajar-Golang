package main
import "fmt"

func main(){
	var input int = 20

	for i := 1; i <= input; i++{
		if input%i==0 {
			fmt.Println(i)
		}
	}

}