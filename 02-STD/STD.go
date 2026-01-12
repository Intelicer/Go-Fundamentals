/*
Standard Library - Strings and Sorting

This file demonstrates:
- String manipulation functions (commented out)
- Sorting integer slices
- Sorting string slices
- Searching in sorted slices
*/
package main

import (
	"fmt"
	"sort"
	// "strings" // Uncomment to use string functions
)

func main() {
	workingWithStrings()
	sortingInts()
	sortingStrings()
}

// ============================================
// STRING MANIPULATION (Using strings package)
// ============================================

func workingWithStrings() {
	fmt.Println("\n=== STRING FUNCTIONS ===")

	// Uncomment the strings import at the top to use these

	// greeting := "greetings humans"

	// strings.Contains() - check if substring exists
	// fmt.Println(strings.Contains(greeting, "hello"))     // false
	// fmt.Println(strings.Contains(greeting, "greetings")) // true

	// strings.ReplaceAll() - replace all occurrences
	// fmt.Println(strings.ReplaceAll(greeting, "greetings", "yo"))

	// strings.ToUpper() - convert to uppercase
	// fmt.Println(strings.ToUpper(greeting)) // "GREETINGS HUMANS"

	// strings.Split() - split string into slice
	// fmt.Println(strings.Split(greeting, " ")) // ["greetings", "humans"]

	fmt.Println("(String functions are commented out)")
}

// ============================================
// SORTING INTEGERS
// ============================================

func sortingInts() {
	fmt.Println("\n=== INT SORTING ===")

	// Create an unsorted slice of integers
	ages := []int{45, 20, 30, 56, 49, 24, 59}

	fmt.Println("Before sorting:", ages)

	// sort.Ints() sorts the slice in ascending order
	// Important: This modifies the original slice (in-place sorting)
	sort.Ints(ages)

	fmt.Println("After sorting:", ages)

	fmt.Println("\n--- Searching in Sorted Slice ---")

	// sort.SearchInts() returns the index where value is (or should be)
	// Important: Only works correctly on SORTED slices

	index1 := sort.SearchInts(ages, 30)
	fmt.Printf("Index of 30: %v\n", index1)

	// If element doesn't exist, returns where it WOULD be inserted
	index2 := sort.SearchInts(ages, 100)
	fmt.Printf("Index for 100 (not in slice): %v (length: %v)\n", index2, len(ages))

	// Example: searching for 25 (not in slice)
	index3 := sort.SearchInts(ages, 25)
	fmt.Printf("Index for 25 (not in slice): %v (between 24 and 30)\n", index3)
}

// ============================================
// SORTING STRINGS
// ============================================

func sortingStrings() {
	fmt.Println("\n=== STRING SORTING ===")

	// Create an unsorted slice of strings
	names := []string{"Ahmed", "Raghd", "Hassan", "Mahmoud", "Nader"}

	fmt.Println("Before sorting:", names)

	// sort.Strings() sorts alphabetically (case-sensitive)
	// Uppercase letters come before lowercase in sorting
	sort.Strings(names)

	fmt.Println("After sorting:", names)

	fmt.Println("\n--- Searching in Sorted String Slice ---")

	// sort.SearchStrings() works like SearchInts but for strings
	index := sort.SearchStrings(names, "Raghd")
	fmt.Printf("Index of 'Raghd': %v\n", index)

	// Search for non-existent string
	index2 := sort.SearchStrings(names, "Zara")
	fmt.Printf("Index for 'Zara' (not in slice): %v\n", index2)
}
