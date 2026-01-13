package main

import (
    "13-ReciverFunctionsWithAndWithoutPointers/structs"
    "fmt"
)

func main() {
    // =====================
    // USING RECEIVER FUNCTIONS (METHODS)
    // =====================

    // Create a map with extra information
    information := map[string]string{
        "Location":  "Amesterdam",
        "HouseType": "Appartment",
    }

    // Create a new Person using the constructor
    person := structs.NewPerson("Mahmoud", 20, information)

    // Call a method on the person struct
    // This prints all the person's info in a formatted way
    fmt.Println("\n" + person.PersonFormattedInformation())

    // =====================
    // MODIFYING WITH POINTER RECEIVER
    // =====================

    // UpdateAge uses a pointer receiver (*Person)
    // So it modifies the ORIGINAL person, not a copy
    person.UpdateAge(10)
    fmt.Println("\n" + person.PersonFormattedInformation())

    // =====================
    // ADDING MORE DATA
    // =====================

    // AddExtraInformation also uses a pointer receiver
    // The new info is added to the original person
    person.AddExtraInformation("HouseNumber", "20A")
    fmt.Println("\n" + person.PersonFormattedInformation())
}