package main

import (
    "fmt"
    "12-StructsCustomType/structs"
)

func main() {
    // =====================
    // USING STRUCTS
    // =====================

    // Create a new Person using the constructor function
    // NewPerson returns a Person struct with Name and Age set
    person := structs.NewPerson("Mahmoud", 20)

    // Access struct fields using dot notation
    // %s = string, %d = integer
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}