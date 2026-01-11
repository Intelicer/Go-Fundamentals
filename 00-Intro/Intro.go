// Package main groups basic type demos.
package main

import "fmt"

func main() {
	// Uncomment to run individual demos.
	// workingWithString()
	// workingWithInt()
	// workingWithFloat()
}

func workingWithString() {
	// Type comes after the variable name.
	var nameOne string = "mahmoud" // Explicit declaration.
	var nameTwo = "madhun"         // Implicit declaration.

	// Cannot leave an uninitialized variable without a declared type.
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "S"

	fmt.Println("firstName: " + nameOne)
	fmt.Println("lastName: " + nameTwo)
	fmt.Println("Extra text: " + nameThree)

	/*
		Short variable declaration is only available inside functions.
		It is a shorthand for var within local scope.
	*/
	firstNum := "testing"

	fmt.Println(firstNum)

	// Formatted strings.
	fmt.Printf("firstname %v and lastname %v \n", nameOne, nameTwo) // %v inserts values in order.
	fmt.Printf("firstname %q and lastname %q \n", nameOne, nameTwo) // %q wraps values in quotes.
	fmt.Printf("The given variable is of type: %T \n", nameOne)     // %T shows the variable type.

	str := fmt.Sprintf("firstname %v and lastname %v \n", nameOne, nameTwo) // Store formatted output.

	fmt.Println("text: ", str)
}

func workingWithInt() {
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 54

	fmt.Println(ageOne, ageTwo, ageThree)
}

// Floats have two common bit sizes: 32 or 64.
func workingWithFloat() {

	// var weight float32 = 123.3
	// var weight2 float64 = 17187.23
	// fmt.Println(weight, weight2)
	// fmt.Print(weight, weight2, "\n") // Use \n to force a newline when printing.
	// fmt.Print("test")

	// Formatted float output.
	fmt.Printf("you scored %f points!\n", 203.23)
	fmt.Printf("you scored %0.1f points!\n", 203.23)

}

/*
* bits & memory
* https://go.dev/ref/spec#Numeric_types
 */
func bitsAndMem() {
	// var numOne int8 = 215 // Out of range for int8.
}
