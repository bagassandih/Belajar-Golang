package main
import "fmt"

func main(){
	var studentScore int = 32

	switch{
	case studentScore >= 80:
	fmt.Println(" A ")
	case studentScore >= 65:
	fmt.Println(" B+ ")
	case studentScore >= 50:
	fmt.Println(" B ")
	case studentScore >= 35:
	fmt.Println(" C ")
	case studentScore >= 0:
	fmt.Println(" D ")
	}

}