/*
Loops in Go

This file demonstrates:
- Traditional for loop with initialization, condition, and increment
- While-style loop (condition only)
- Iterating over slices with index
- Range-based loops (with and without index)

Note: Go does NOT have a while loop - use for with condition only
*/
package main

import "fmt"

func main() {
	traditionalForLoop()
	whileStyleLoop()
	iteratingWithIndex()
	rangeWithIndexAndValue()
	rangeWithValueOnly()
}

// ============================================
// TRADITIONAL FOR LOOP
// ============================================

// traditionalForLoop demonstrates classic C-style for loop
func traditionalForLoop() {
	fmt.Println("\n=== TRADITIONAL FOR LOOP ===")

	// Syntax: for initialization; condition; increment
	// - i := 1        : Initialize counter
	// - i <= 5        : Continue while this is true
	// - i++           : Increment after each iteration
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop 1 iteration:", i)
	}
}

// ============================================
// WHILE-STYLE LOOP
// ============================================

// whileStyleLoop shows how to create a while-loop in Go
func whileStyleLoop() {
	fmt.Println("\n=== WHILE-STYLE LOOP ===")

	// Go doesn't have "while" keyword
	// Use "for" with only a condition (acts like while)
	x := 1

	for x <= 5 {
		fmt.Println("Loop 2 iteration:", x)
		x++ // Must manually increment
	}

	// Infinite loop example (commented out):
	// for {
	//     fmt.Println("This runs forever!")
	//     break // Use break to exit
	// }
}

// ============================================
// ITERATING WITH INDEX
// ============================================

// iteratingWithIndex demonstrates traditional index-based iteration
func iteratingWithIndex() {
	fmt.Println("\n=== ITERATING WITH INDEX ===")

	names := []string{"Mahmoud", "Ahmed", "Raghd", "Rezan", "Mohammed"}

	// Classic approach: use index to access elements
	// len() returns the length of the slice
	for i := 0; i < len(names); i++ {
		fmt.Printf("Student [%v]: %v\n", i, names[i])
	}
}

// ============================================
// RANGE WITH INDEX AND VALUE
// ============================================

// rangeWithIndexAndValue shows the "range" keyword for easier iteration
func rangeWithIndexAndValue() {
	fmt.Println("\n=== RANGE WITH INDEX AND VALUE ===")

	names := []string{"Mahmoud", "Ahmed", "Raghd", "Rezan", "Mohammed"}

	// "range" keyword returns both index and value
	// This is the preferred way to iterate in Go
	for index, value := range names {
		fmt.Printf("Position %v: %v\n", index, value)

		// You can modify the slice using the index
		// names[index] = "Modified"
	}

	// Note: "value" is a COPY, modifying it doesn't change the slice
	// To modify the slice, use: names[index] = newValue
}

// ============================================
// RANGE WITH VALUE ONLY
// ============================================

// rangeWithValueOnly shows how to ignore the index using underscore
func rangeWithValueOnly() {
	fmt.Println("\n=== RANGE WITH VALUE ONLY ===")

	names := []string{"Mahmoud", "Ahmed", "Raghd", "Rezan", "Mohammed"}

	// Use underscore "_" to ignore the index
	// The blank identifier "_" discards values you don't need
	for _, value := range names {
		fmt.Println("Name:", value)
	}

	// You can also ignore the value and keep only the index:
	// for index := range names {
	//     fmt.Println("Index:", index)
	// }
}
