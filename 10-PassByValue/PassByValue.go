package main

import "fmt"

func main() {
	name := "mahmoud"

	updateName(&name)

	name = updateName2(name)

	fmt.Println(name)



	menu := map[string]float64{
		"pie":55.0,
		"ice":14.2,
	}

	updateMenu(menu, "coffee", 2.1)

	for key,value := range menu {
		fmt.Println(key, value)
	}

}

//solution1

//adding pointer as parameter and then use the reference of the name we want to change
func updateName(name *string){
	*name = "run"
}

//solution2

func updateName2(name string) string{
	name = "testing out"
	return name
}

func updateMenu(menuToChange map[string]float64, productName string, price float64){
	menuToChange[productName] = price
}
