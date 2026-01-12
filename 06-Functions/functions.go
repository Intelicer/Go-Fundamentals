/*
Functions in Go

This file demonstrates:
- Basic function definitions
- Functions with parameters
- Functions with return values
- Functions as parameters (higher-order functions)
- Working with the math package
- Formatted output with Printf
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	basicFunctions()
	functionWithReturn()
	higherOrderFunction()
}

// ============================================
// BASIC FUNCTIONS
// ============================================

// basicFunctions demonstrates simple function calls
func basicFunctions() {
	fmt.Println("\n=== BASIC FUNCTIONS ===")

	// Call functions with parameters
	sayGreetings("Mahmoud")
	sayHaveNiceDay("Ahmed")
}

// sayGreetings is a simple function that takes a string parameter
// - "name string" declares a parameter named "name" of type string
// - No return type specified (void function)
func sayGreetings(name string) {
	fmt.Printf("Good morning %v\n", name)
}

// sayHaveNiceDay shows another simple function example
func sayHaveNiceDay(name string) {
	fmt.Printf("Have a nice day %v\n", name)
}

// ============================================
// FUNCTIONS WITH RETURN VALUES
// ============================================

// functionWithReturn demonstrates functions that return values
func functionWithReturn() {
	fmt.Println("\n=== FUNCTIONS WITH RETURN VALUES ===")

	// Call function and store the returned value
	area := circleArea(10.5)

	// Printf with format specifier:
	// %0.3f = floating-point with 3 decimal places
	fmt.Printf("Circle area = %0.3f\n", area)

	// More examples
	area2 := circleArea(5.0)
	fmt.Printf("Circle with radius 5.0 = %0.2f\n", area2)
}

// circleArea calculates the area of a circle given its radius
// - "radius float64" = parameter
// - "float64" after () = return type
// - Must use "return" keyword to send value back
func circleArea(radius float64) float64 {
	// math.Pi = constant from math package (3.141592...)
	// math.Pow(radius, 2) = radius squared (radius²)
	// Formula: Area = π × r²
	return math.Pi * math.Pow(radius, 2)
}

// ============================================
// HIGHER-ORDER FUNCTIONS
// ============================================

// higherOrderFunction demonstrates passing functions as parameters
func higherOrderFunction() {
	fmt.Println("\n=== HIGHER-ORDER FUNCTIONS ===")

	names := []string{"Raghd", "Mahmoud", "Ahmed", "Rezan"}

	// Pass the sayGreetings function as a parameter
	// Notice: no () after sayGreetings (we pass the function itself)
	cycleNames(names, sayGreetings)

	fmt.Println()

	// Pass a different function
	cycleNames(names, sayHaveNiceDay)
}

// cycleNames is a higher-order function
// It takes a slice and another function as parameters
//   - "names []string" = slice of strings
//   - "functionToCall func(string)" = parameter that is a function
//     The function must take one string parameter and return nothing
func cycleNames(names []string, functionToCall func(string)) {

	// Loop through all names
	for _, name := range names {
		// Call the function that was passed as parameter
		functionToCall(name)
	}
}

// Benefits of higher-order functions:
// 1. Code reuse - cycleNames works with ANY function that matches the signature
// 2. Flexibility - behavior can be changed without modifying cycleNames
// 3. Functional programming - enables powerful patterns
