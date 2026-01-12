/*
Multiple Return Values in Go

This file demonstrates:
- Functions that return multiple values
- Unpacking multiple return values
- Using the strings package for text manipulation
- String slicing syntax
- Handling edge cases with conditional returns
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	multipleReturnExample()
	handlingEdgeCases()
}

// ============================================
// MULTIPLE RETURN VALUES
// ============================================

// multipleReturnExample demonstrates unpacking multiple returns
func multipleReturnExample() {
	fmt.Println("\n=== MULTIPLE RETURN VALUES ===")

	// Call function that returns TWO values
	// Use two variables to "catch" both returns
	firstName, lastName := getInitials("Mahmoud Ahmed")

	fmt.Printf("First initial: %v\n", firstName)
	fmt.Printf("Last initial: %v\n", lastName)

	// More examples
	first, last := getInitials("John Doe")
	fmt.Printf("%v.%v.\n", first, last) // J.D.

	// If you only need one value, use underscore to ignore the other
	first, _ = getInitials("Alice Wonderland")
	fmt.Println("First initial only:", first)
}

// ============================================
// HANDLING EDGE CASES
// ============================================

// handlingEdgeCases shows how the function handles single names
func handlingEdgeCases() {
	fmt.Println("\n=== EDGE CASES ===")

	// Single name (no space)
	first, last := getInitials("Madonna")
	fmt.Printf("Single name: %v %v\n", first, last) // M _

	// Multiple names
	first2, last2 := getInitials("John Paul Smith")
	fmt.Printf("Multiple names (takes first 2): %v %v\n", first2, last2) // J P
}

// ============================================
// FUNCTION WITH MULTIPLE RETURNS
// ============================================

// getInitials extracts the first letter of each word
// Returns: (firstInitial, lastInitial)
// - "(string, string)" after parameters = returns TWO strings
func getInitials(sentence string) (string, string) {

	// Step 1: Split the sentence into words
	// strings.Split(sentence, " ") splits by space character
	// Example: "Mahmoud Ahmed" → ["Mahmoud", "Ahmed"]
	nameSlice := strings.Split(sentence, " ")

	// Step 2: Create empty slice to store initials
	var initials []string

	// Step 3: Loop through each word and extract first character
	for _, value := range nameSlice {

		// String slicing syntax: value[start:end]
		// value[:1] = from beginning to index 1 (first character)
		// Example: "Mahmoud"[:1] = "M"

		// Commented code shows alternative approach:
		// newValue := value[:1] + value[1:]  // Would give full word
		// initials = append(initials, newValue)

		// Append just the first letter
		initials = append(initials, value[:1])
	}

	// Step 4: Return based on how many words we found

	// If 2 or more words, return first two initials
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	// If only 1 word, return initial and underscore as placeholder
	// This is a default/fallback value
	return initials[0], "_"
}
