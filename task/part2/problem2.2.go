package main
import "fmt"

func main(){
	var bilangan int = 20

	for i := bilangan; i >= 1; i--{
		if bilangan%i==0 {
			fmt.Println(bilangan, i)
		}
	}

}