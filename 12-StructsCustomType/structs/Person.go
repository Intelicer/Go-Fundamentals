package structs

// =====================
// STRUCTS IN GO
// =====================

// A struct is a custom type that groups related data together
// Similar to a class in other languages (but without methods here)
type Person struct {
    Name string // Person's name
    Age  int    // Person's age
}

// =====================
// CONSTRUCTOR FUNCTION
// =====================

// Go doesn't have built-in constructors
// Instead, we create a function that returns a new struct
// Convention: Name it "New" + struct name (e.g., NewPerson)
func NewPerson(name string, age int) Person {
    person := Person{
        Name: name,
        Age:  age,
    }
    return person
}

// =====================
// QUICK REFERENCE
// =====================
// type Name struct { }  -> define a struct
// Name{field: value}    -> create a struct instance
// instance.Field        -> access a field