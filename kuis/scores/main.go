package main

import "fmt"

func main(){

	/*
	Cek value dari array scores, jika value lebih dari 90 akan masuk ke slice goodScores
	*/

	scores := [8]int{100, 88, 75, 92, 70, 93, 88, 67}

	var goodScores []int

	for _, daftarScores := range scores {
		if daftarScores >= 90 {
			goodScores = append(goodScores, daftarScores)
		} 

	}

	fmt.Println(goodScores)

}