package main

import (
    "14-UserInput/structs"
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // =====================
    // USER INPUT EXAMPLE
    // =====================

    // Create a person by asking the user for input
    person := createPerson()

    // Display the person's formatted information
    fmt.Println("\n" + person.PersonFormattedInformation())
}

// =====================
// CREATE PERSON FROM USER INPUT
// =====================

// createPerson asks the user for name, age, and optional extra info
// Returns a fully constructed Person struct
func createPerson() structs.Person {
    // Control flags for loops
    loopAddInformation := true
    loopAddExtraInfo := false

    // bufio.Reader reads input from the terminal
    // os.Stdin = standard input (keyboard)
    reader := bufio.NewReader(os.Stdin)

    // =====================
    // GET NAME
    // =====================
    fmt.Println("\nPerson name? ")
    name, _ := reader.ReadString('\n')  // Read until user presses Enter
    name = strings.TrimSpace(name)       // Remove extra whitespace/newline

    // =====================
    // GET AGE
    // =====================
    fmt.Println("\nPerson age? ")
    ageStr, _ := reader.ReadString('\n')
    ageStr = strings.TrimSpace(ageStr)
    age, _ := strconv.Atoi(ageStr)       // Convert string to int

    // =====================
    // ASK FOR EXTRA INFO (OPTIONAL)
    // =====================
    fmt.Println("\nDo you want to add personal information? (optional) use either yes or no ")

    // Loop until user gives a valid answer (yes/no)
    for loopAddInformation {
        userLoopAnswer, _ := reader.ReadString('\n')
        userLoopAnswer = strings.TrimSpace(userLoopAnswer)
        userLoopAnswer = strings.ToUpper(userLoopAnswer) // Make uppercase for comparison

        switch userLoopAnswer {
        case "YES":
            loopAddExtraInfo = true      // User wants to add extra info
            loopAddInformation = false   // Exit this loop
        case "NO":
            loopAddInformation = false   // Exit this loop
        default:
            fmt.Println("\nRewrite again Yes or No.")
        }
    }

    // =====================
    // COLLECT EXTRA INFORMATION
    // =====================

    // Empty map to store extra info
    information := map[string]string{}

    // Loop to add multiple pieces of extra info
    for loopAddExtraInfo {
        // Get the type of info (e.g., "Location", "Email")
        fmt.Println("\nWhat type of info?")
        typeOfInfo, _ := reader.ReadString('\n')
        typeOfInfo = strings.TrimSpace(typeOfInfo)

        // Get the details (e.g., "Amsterdam", "test@test.com")
        fmt.Println("\nWrite Detail about info")
        detailsOfInfo, _ := reader.ReadString('\n')
        detailsOfInfo = strings.TrimSpace(detailsOfInfo)

        // Add to the map
        information[typeOfInfo] = detailsOfInfo

        // Ask if user wants to add more
        fmt.Println("\nDo you want add more info? answer by yes or no")
        addMoreInfo, _ := reader.ReadString('\n')
        addMoreInfo = strings.TrimSpace(addMoreInfo)
        addMoreInfo = strings.ToUpper(addMoreInfo)

        // Validate yes/no answer
        for loopAddExtraInfo {
            switch addMoreInfo {
            case "NO":
                loopAddExtraInfo = false // Stop adding info
            case "YES":
                fmt.Println("\nYou will be getting same questions to add more information.")
            default:
                fmt.Println("\nYou must answer by Yes or No")
            }
        }
    }

    fmt.Println("\nPerson obj made.")

    // Return the new Person using the constructor
    return structs.NewPerson(name, age, information)
}

// =====================
// QUICK REFERENCE
// =====================
// bufio.NewReader(os.Stdin)  -> create a reader for keyboard input
// reader.ReadString('\n')    -> read until Enter is pressed
// strings.TrimSpace(s)       -> remove whitespace and newlines
// strings.ToUpper(s)         -> convert to uppercase
// strconv.Atoi(s)            -> convert string to int