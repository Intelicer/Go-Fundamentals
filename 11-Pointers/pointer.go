package main

import "fmt"

func main() {
    // =====================
    // POINTERS IN GO
    // =====================

    name := "mahmoud"

    // & gives you the memory address of a variable
    fmt.Println(&name) // Prints something like: 0xc0000140a0

    // This does NOT change the original - Go passes a COPY
    updateName(name)

    // This DOES change the original - we pass the memory address
    updateNameSolution(&name)

    // %v = value, %q = quoted string
    fmt.Printf("\nAddress (%v) has value %q", &name, name)
    // Output: Address (0xc0000140a0) has value "raghd"
}

// =====================
// WITHOUT POINTER (COPY)
// =====================

// This receives a COPY of the value
// Changes here do NOT affect the original variable
func updateName(name string) {
    name = "run"                // Only changes the local copy
    fmt.Println(&name)          // Different address than the original
}

// =====================
// WITH POINTER (REFERENCE)
// =====================

// *string means "pointer to a string"
// This receives the MEMORY ADDRESS of the original variable
func updateNameSolution(name *string) {
    // *name means "the value at this address"
    // This DOES change the original variable
    *name = "raghd"
}

// =====================
// QUICK REFERENCE
// =====================
// &variable  -> get the memory address
// *pointer   -> get/set the value at that address
// *type      -> declare a pointer type (e.g., *string)