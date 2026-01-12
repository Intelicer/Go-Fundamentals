package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup": 4.18,
		"rice": 1.98,
	}

	// fmt.Println(menu)

	fmt.Println(menu["rice"])

	//map loop
	for key,value := range menu {
		fmt.Println(key, value)
	}


	// ints as key type

	phonebook := map[int]string{
		2675884967 : "Random",
		96700651 : "Mahmoud",
		83828492 : "Idiot",
	}
	// for key,value := range phonebook {
	// 	fmt.Println(key, value)
	// }
	
	fmt.Println(phonebook[96700651])
	
	phonebook[96700651] = "non of your business"
	
	fmt.Println(phonebook[96700651])
}