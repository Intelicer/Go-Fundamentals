package main

import (
  "13-ReciverFunctionsWithAndWithoutPointers/structs"
	"fmt"
)

func main() {
	information := map[string]string{
		"Location":  "Amesterdam",
		"HouseType": "Appartment",
	}
	person := structs.NewPerson("Mahmoud", 20, information)
	fmt.Println("\n" + person.PersonFormattedInformation())
	
	person.UpdateAge(10)
	fmt.Println("\n" + person.PersonFormattedInformation())
	
	
	person.AddExtraInformation("HouseNumber","20A")
	fmt.Println("\n" + person.PersonFormattedInformation())
	
}
