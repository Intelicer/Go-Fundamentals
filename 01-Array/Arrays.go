/*
Arrays and Slices in Go

This file demonstrates:
- Fixed-size arrays
- Dynamic slices
- Appending to slices
- Slice ranges and subslicing
*/
package main

import "fmt"

func main() {
	dealingWithArray()
	dealingWithSlices()
}

// ============================================
// ARRAYS - Fixed Size Collections
// ============================================

// dealingWithArray demonstrates fixed-size array operations
func dealingWithArray() {
	fmt.Println("\n=== ARRAYS ===")

	// Arrays have FIXED size defined at declaration
	// Method 1: Explicit type and size
	// var ages [3]int = [3]int{20, 25, 30}

	// Method 2: Implicit type (shorter syntax)
	// var agesTwo = [3]int{20, 25, 30}

	// Method 3: Short declaration (most common)
	names := [3]string{"mario", "max", "mose"}

	// Print array and its length
	fmt.Println("Original names:", names, "Length:", len(names))

	// Access and modify array elements by index (0-based)
	names[0] = "test"
	fmt.Println("After modification:", names, "Length:", len(names))

	// Important: Array size is fixed - you cannot add or remove elements
	// If you need dynamic size, use slices instead
}

// ============================================
// SLICES - Dynamic Collections
// ============================================

// dealingWithSlices demonstrates dynamic slice operations
func dealingWithSlices() {
	fmt.Println("\n=== SLICES ===")

	// Slices are dynamic - no fixed size in declaration
	// Notice: []int instead of [3]int
	scores := []int{100, 50, 20}

	fmt.Println("Initial scores:", scores, "Length:", len(scores))

	// append() adds elements to a slice
	// Important: append returns a NEW slice, must reassign
	scores = append(scores, 120)

	fmt.Println("After append:", scores, "Length:", len(scores))

	fmt.Println("\n--- Slice Ranges (Subslicing) ---")

	// Slice ranges use syntax: slice[start:end]
	// - start: inclusive
	// - end: exclusive

	// Get elements from index 1 to 2 (3 is excluded)
	// middleScores := scores[1:3]

	// Get all elements from index 1 to the end
	middleScores := scores[1:]
	fmt.Println("Middle scores [1:]:", middleScores)

	// Other range examples:
	// scores[:2]  = from start to index 1 (2 excluded)
	// scores[:]   = entire slice (copy)
	// scores[1:3] = from index 1 to 2 (3 excluded)

	// Additional slice operations
	fmt.Println("\n--- More Slice Operations ---")

	// Get first two elements
	firstTwo := scores[:2]
	fmt.Println("First two:", firstTwo)

	// Get last two elements
	lastTwo := scores[len(scores)-2:]
	fmt.Println("Last two:", lastTwo)
}
