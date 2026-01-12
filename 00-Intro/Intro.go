/*
Package main - Introduction to Go Basics

This file covers fundamental Go concepts:
- Variable declarations (explicit, implicit, short)
- Basic data types (string, int, float)
- String formatting and output
- Type information
*/
package main

import "fmt"

func main() {
	// Uncomment the function you want to run
	workingWithString()
	// workingWithInt()
	// workingWithFloat()
	// bitsAndMem()
}

// ============================================
// STRING DEMONSTRATIONS
// ============================================

// workingWithString demonstrates different ways to declare and use strings in Go
func workingWithString() {

	// Method 1: Explicit type declaration with initialization
	var nameOne string = "mahmoud"

	// Method 2: Implicit type declaration (Go infers the type)
	var nameTwo = "madhun"

	// Method 3: Declaration without initialization (defaults to empty string "")
	var nameThree string

	// Print all three variables
	fmt.Println(nameOne, nameTwo, nameThree)

	// Assign value to the previously empty variable
	nameThree = "S"

	// String concatenation using the + operator
	fmt.Println("firstName: " + nameOne)
	fmt.Println("lastName: " + nameTwo)
	fmt.Println("Extra text: " + nameThree)

	// Method 4: Short variable declaration (only works inside functions)
	// This is the most common way to declare variables in Go
	firstNum := "testing"
	fmt.Println(firstNum)

	fmt.Println("\n--- Formatted String Output ---")

	// Printf allows formatted output with placeholders
	// %v = value in default format
	fmt.Printf("firstname %v and lastname %v \n", nameOne, nameTwo)

	// %q = value wrapped in double quotes
	fmt.Printf("firstname %q and lastname %q \n", nameOne, nameTwo)

	// %T = shows the data type of the variable
	fmt.Printf("The given variable is of type: %T \n", nameOne)

	// Sprintf returns the formatted string instead of printing it
	// Useful when you need to store or manipulate the formatted output
	str := fmt.Sprintf("firstname %v and lastname %v \n", nameOne, nameTwo)
	fmt.Println("Stored text: ", str)
}

// ============================================
// INTEGER DEMONSTRATIONS
// ============================================

// workingWithInt shows different ways to declare integer variables
func workingWithInt() {

	// Three ways to declare integers (same as strings)
	var ageOne int = 20 // Explicit
	var ageTwo = 30     // Implicit
	ageThree := 54      // Short declaration

	fmt.Println(ageOne, ageTwo, ageThree)

	// Additional integer operations
	fmt.Printf("Type of ageOne: %T\n", ageOne)
	fmt.Printf("Sum: %v\n", ageOne+ageTwo+ageThree)
}

// ============================================
// FLOAT DEMONSTRATIONS
// ============================================

// workingWithFloat demonstrates floating-point numbers and precision formatting
func workingWithFloat() {

	// Float types: float32 (32-bit) and float64 (64-bit, more precise)
	var weight float32 = 123.3
	var weight2 float64 = 17187.23

	fmt.Println(weight, weight2)

	// Print vs Println: Print requires manual newline character \n
	fmt.Print(weight, weight2, "\n")
	fmt.Print("test\n")

	fmt.Println("\n--- Float Formatting ---")

	// %f = floating-point number (default 6 decimal places)
	fmt.Printf("You scored %f points!\n", 203.23)

	// %0.1f = floating-point with 1 decimal place
	fmt.Printf("You scored %0.1f points!\n", 203.23)

	// %0.2f = 2 decimal places
	fmt.Printf("You scored %0.2f points!\n", 203.23)
}

// ============================================
// NUMERIC TYPE INFORMATION
// ============================================

/*
   bitsAndMem provides information about numeric type ranges

   Integer types and their ranges:
   - int8:   -128 to 127
   - int16:  -32,768 to 32,767
   - int32:  -2,147,483,648 to 2,147,483,647
   - int64:  -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
   - uint8:  0 to 255 (unsigned, positive only)
   - uint16: 0 to 65,535
   - uint32: 0 to 4,294,967,295
   - uint64: 0 to 18,446,744,073,709,551,615

   Float types:
   - float32: ~6-7 decimal digits precision
   - float64: ~15-16 decimal digits precision (recommended)

   Reference: https://go.dev/ref/spec#Numeric_types
*/
func bitsAndMem() {
	// This would cause a compile error - 215 exceeds int8 range (-128 to 127)
	// var numOne int8 = 215

	// Correct usage:
	var numOne int8 = 127 // Maximum value for int8
	fmt.Printf("Max int8: %v\n", numOne)

	var numTwo uint8 = 255 // Maximum value for uint8
	fmt.Printf("Max uint8: %v\n", numTwo)
}
