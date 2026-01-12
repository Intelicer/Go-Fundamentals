/*
Conditional Expressions and Control Flow

This file demonstrates:
- if, else if, else statements
- Logical AND (&&) operator
- Logical OR (||) operator
- Loop control with continue statement
- Combining loops with conditionals
*/
package main

import "fmt"

func main() {
	ifElseStatements()
	logicalAndOperator()
	logicalOrOperator()
	loopWithContinue()
}

// ============================================
// IF-ELSE STATEMENTS
// ============================================

// ifElseStatements demonstrates basic conditional logic
func ifElseStatements() {
	fmt.Println("\n=== IF-ELSE STATEMENTS ===")

	age := 23

	// if-else chain: checks conditions top to bottom
	// Executes first matching condition, then skips the rest
	if age > 30 {
		// This block runs if age is greater than 30
		fmt.Println("Age is:", age)
	} else if age < 20 {
		// This block runs if age is less than 20
		fmt.Println("You are young")
	} else {
		// This block runs if none of the above conditions match
		fmt.Println("Age is between 20-30:", age)
	}
}

// ============================================
// LOGICAL AND OPERATOR (&&)
// ============================================

// logicalAndOperator shows how to combine multiple conditions with AND
func logicalAndOperator() {
	fmt.Println("\n=== LOGICAL AND (&&) ===")

	age := 23

	// AND operator (&&): BOTH conditions must be true
	// true && true = true
	// true && false = false
	// false && true = false
	// false && false = false
	if age > 20 && age < 30 {
		fmt.Println("Age is between 20 and 30")
	}

	// More examples of AND
	hasID := true
	isAdult := age >= 18

	if hasID && isAdult {
		fmt.Println("Can enter the venue")
	}

	// Multiple AND conditions
	if age > 20 && age < 30 && hasID {
		fmt.Println("All three conditions are true")
	}
}

// ============================================
// LOGICAL OR OPERATOR (||)
// ============================================

// logicalOrOperator demonstrates the OR operator
func logicalOrOperator() {
	fmt.Println("\n=== LOGICAL OR (||) ===")

	age := 23

	// OR operator (||): AT LEAST ONE condition must be true
	// true || true = true
	// true || false = true
	// false || true = true
	// false || false = false
	if age > 20 || age < 30 {
		fmt.Println("Age is greater than 20 OR less than 30")
		fmt.Println("(In this case, age 23 satisfies BOTH conditions)")
	}

	// Practical OR example
	isWeekend := false
	isHoliday := true

	if isWeekend || isHoliday {
		fmt.Println("You can rest today!")
	}

	// Combining AND and OR
	hasTicket := true
	if (age >= 18 || hasTicket) && true {
		fmt.Println("Complex condition example")
	}
}

// ============================================
// LOOP WITH CONTINUE
// ============================================

// loopWithContinue demonstrates loop control flow
func loopWithContinue() {
	fmt.Println("\n=== LOOP WITH CONTINUE ===")

	names := []string{"mario", "luigi", "peach"}

	// range loop with conditional logic
	for index, value := range names {

		// continue: skip the rest of THIS iteration and go to the next
		// break: exit the entire loop
		if index == 1 {
			fmt.Printf("\nSkipping index %v (value: %v)\n", index, value)
			continue // Skip printing "luigi"
		}

		// This line is skipped when continue is executed
		fmt.Printf("Position %v: %v\n", index, value)
	}

	fmt.Println("\n--- Break Example ---")

	// break example: stop the loop entirely
	for index, value := range names {
		if index == 1 {
			fmt.Println("Breaking at index 1")
			break // Exit the loop completely
		}
		fmt.Printf("Position %v: %v\n", index, value)
	}
}
