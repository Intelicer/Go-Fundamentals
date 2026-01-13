package main

import (
    "fmt"
    "12-StructsCustomType/structs"
)

func main() {
    p := structs.NewPerson("Mahmoud", 20)
    fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}