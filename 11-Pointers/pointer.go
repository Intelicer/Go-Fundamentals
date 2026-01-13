package main

import "fmt"

func main() {
	name := "mahmoud"



	fmt.Println(&name) //address of the variable we want to use


	updateName(name)

	updateNameSolution(&name)

	fmt.Printf("\nAddress (%v) has value %q",&name, name)
}


func updateName(name string){
	name = "run"
	fmt.Println(&name) // address of the variable in the scope
}





//solution1

//adding pointer as parameter and then use the reference of the name we want to change
func updateNameSolution(name *string){
	*name = "raghd"
}

