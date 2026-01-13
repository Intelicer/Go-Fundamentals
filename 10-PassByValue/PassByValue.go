package main

import "fmt"

func main() {
    // =====================
    // PASS BY VALUE VS POINTERS
    // =====================

    name := "mahmoud"

    // Solution 1: Use a pointer to modify the original variable
    updateName(&name) // Pass the memory address of 'name'

    // Solution 2: Return a new value and reassign it
    name = updateName2(name)

    fmt.Println(name) // Output: testing out

    // =====================
    // MAPS ARE REFERENCE TYPES
    // =====================

    // Maps are passed by reference automatically
    // Changes inside a function affect the original map
    menu := map[string]float64{
        "pie": 55.0,
        "ice": 14.2,
    }

    // No need to use pointers - the map is modified directly
    updateMenu(menu, "coffee", 2.1)

    // Print all menu items (now includes "coffee")
    for key, value := range menu {
        fmt.Println(key, value)
    }
}

// =====================
// SOLUTION 1: POINTERS
// =====================

// Use a pointer (*string) to modify the original variable
// The * means "the value at this address"
func updateName(name *string) {
    *name = "run" // Dereference and update the original
}

// =====================
// SOLUTION 2: RETURN VALUE
// =====================

// Pass by value (copy), modify it, and return the new value
func updateName2(name string) string {
    name = "testing out"
    return name // Caller must reassign: name = updateName2(name)
}

// =====================
// MAPS (REFERENCE TYPE)
// =====================

// Maps don't need pointers - they're reference types by default
func updateMenu(menuToChange map[string]float64, productName string, price float64) {
    menuToChange[productName] = price // Modifies the original map
}