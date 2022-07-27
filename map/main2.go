package main

import (
	"fmt"
	"reflect"
	)

func main() {
	
//	map[tipedata_key]tipedata_value{array}

//	deklarasi map harus pakai {} diakhir agar nilai default terbaca
	var contoh = map[string]int{}

	contoh["iniKey"] = 1
	contoh["iniKeyJuga"] = 1
	contoh["iniKeyBro"] = 1

// bisa ditimpa
	contoh["iniKey"] = 2
	fmt.Println(contoh)
	fmt.Println(contoh["iniKey"])
	fmt.Println(reflect.ValueOf(contoh).Kind())
	fmt.Println()


// contoh lain
	mobaGame := map[string]string{}

	mobaGame["Dota 2"] = "Valve"
	mobaGame["Lol"] = "Riot"
	mobaGame["ML"] = "Moonton"
	 
	fmt.Println(mobaGame)
// menggunakan for
	for key, value:= range mobaGame {
		fmt.Println(key, "pabriknya", value )
	}

// cek  map berdasarkan key
	cek := mobaGame["Dota 2"] //cek valuenya
	fmt.Println(cek) 
// map
	_, isThere := mobaGame["PUBG"]
	fmt.Println(isThere)
}