package main

import "fmt"

func main() {
    // Creating a map with string keys and float64 values
    // Syntax: map[keyType]valueType
    menu := map[string]float64{
        "soup": 4.18,
        "rice": 1.98,
    }

    // Access a value by its key
    fmt.Println(menu["rice"]) // Output: 1.98

    // Loop through a map using range
    // Each iteration gives you a key and its value
    for key, value := range menu {
        fmt.Println(key, value)
    }

    // Maps can use different types as keys (here: int keys, string values)
    phonebook := map[int]string{
        2675884967: "Random",
        96700651:   "Mahmoud",
        83828492:   "Idiot",
    }

    // Access a value using an int key
    fmt.Println(phonebook[96700651]) // Output: Mahmoud

    // Update a value by assigning to an existing key
    phonebook[96700651] = "non of your business"

    // Print the updated value
    fmt.Println(phonebook[96700651]) // Output: non of your business
}