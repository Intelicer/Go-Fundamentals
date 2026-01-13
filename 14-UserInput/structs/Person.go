package structs

import (
    "sort"
    "strconv"
)

// =====================
// STRUCT DEFINITION
// =====================

// Person holds basic info and extra details in a map
type Person struct {
    Name        string            // Person's name
    Age         int               // Person's age
    Information map[string]string // Extra info (e.g., "email": "test@test.com")
}

// =====================
// CONSTRUCTOR FUNCTION
// =====================

// NewPerson creates and returns a new Person
// It also sorts the information map alphabetically
func NewPerson(name string, age int, information map[string]string) Person {
    person := Person{
        Name:        name,
        Age:         age,
        Information: information,
    }
    person.sortInformation() // Sort info on creation
    return person
}

// =====================
// RECEIVER FUNCTIONS (METHODS)
// =====================

// PersonFormattedInformation returns a nicely formatted string
// Uses pointer receiver (*Person) for consistency
func (person *Person) PersonFormattedInformation() string {
    // strconv.Itoa converts int to string
    personFormatedString := "Person name: " + person.Name +
        ", Age: " + strconv.Itoa(person.Age) + "\nDetailed Info:"

    // Loop through all extra information
    for key, value := range person.Information {
        personFormatedString += "\n" + key + ": " + value
    }
    return personFormatedString
}

// UpdateAge modifies the person's age
// Uses pointer receiver (*Person) so changes affect the original
func (person *Person) UpdateAge(newAge int) {
    person.Age = newAge
}

// AddExtraInformation adds a new key-value pair to Information
// Uses pointer receiver (*Person) to modify the original map
func (person *Person) AddExtraInformation(infoType string, details string) {
    person.Information[infoType] = details
    person.sortInformation() // Keep info sorted
}

// =====================
// PRIVATE METHOD
// =====================

// sortInformation sorts the Information map alphabetically by key
// Lowercase first letter = private (only accessible in this package)
func (person *Person) sortInformation() {
    // Step 1: Create a slice to hold all keys
    keys := make([]string, 0, len(person.Information))

    // Step 2: Extract all keys from the map
    for key := range person.Information {
        keys = append(keys, key)
    }

    // Step 3: Sort the keys alphabetically
    sort.Strings(keys)

    // Step 4: Rebuild the map in sorted order
    newInfo := make(map[string]string)
    for _, key := range keys {
        newInfo[key] = person.Information[key]
    }

    person.Information = newInfo
}

// =====================
// QUICK REFERENCE
// =====================
// func (p Person) Method()   -> value receiver (works on a copy)
// func (p *Person) Method()  -> pointer receiver (modifies original)
// Lowercase method name      -> private (only this package can use it)
// Uppercase method name      -> public (other packages can use it)