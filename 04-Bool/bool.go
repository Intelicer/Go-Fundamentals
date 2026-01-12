/*
Boolean Type and Comparisons

This file demonstrates:
- Boolean variable declarations
- Comparison operators
- Boolean expressions
- Relational operators (<=, >=, ==, !=, <, >)
*/
package main

import "fmt"

func main() {
	basicBooleans()
	comparisonOperators()
	booleanExpressions()
}

// ============================================
// BASIC BOOLEAN DECLARATIONS
// ============================================

// basicBooleans demonstrates boolean variable usage
func basicBooleans() {
	fmt.Println("\n=== BASIC BOOLEANS ===")

	// Boolean type: true or false (lowercase in Go)
	x := true
	y := false

	fmt.Println("x =", x) // true
	fmt.Println("y =", y) // false

	// Boolean is its own type
	fmt.Printf("Type of x: %T\n", x) // bool
}

// ============================================
// COMPARISON OPERATORS
// ============================================

// comparisonOperators shows different comparison operations
func comparisonOperators() {
	fmt.Println("\n=== COMPARISON OPERATORS ===")

	age := 23

	// Less than or equal to (<=)
	fmt.Printf("%v <= 50: %v\n", age, age <= 50) // true
	fmt.Printf("%v <= 23: %v\n", age, age <= 23) // true
	fmt.Printf("%v <= 10: %v\n", age, age <= 10) // false

	fmt.Println("\n--- All Comparison Operators ---")

	// Equal to
	fmt.Printf("%v == 23: %v\n", age, age == 23) // true

	// Not equal to
	fmt.Printf("%v != 23: %v\n", age, age != 23) // false

	// Less than
	fmt.Printf("%v < 30: %v\n", age, age < 30) // true

	// Greater than
	fmt.Printf("%v > 20: %v\n", age, age > 20) // true

	// Greater than or equal to
	fmt.Printf("%v >= 23: %v\n", age, age >= 23) // true
}

// ============================================
// BOOLEAN EXPRESSIONS
// ============================================

// booleanExpressions demonstrates using booleans in conditions
func booleanExpressions() {
	fmt.Println("\n=== BOOLEAN EXPRESSIONS ===")

	age := 23
	hasLicense := true

	// Storing comparison results
	isAdult := age >= 18
	isTeenager := age >= 13 && age < 20

	fmt.Printf("Is adult (age >= 18): %v\n", isAdult)
	fmt.Printf("Is teenager (13-19): %v\n", isTeenager)

	// Using booleans in if statements
	if isAdult && hasLicense {
		fmt.Println("Can drive a car")
	}

	// Negation with ! operator
	if !isTeenager {
		fmt.Println("Not a teenager")
	}
}
