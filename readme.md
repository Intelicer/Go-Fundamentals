# Go Programming Fundamentals

A comprehensive collection of Go programming examples covering core language concepts. Each module is thoroughly documented with educational comments to facilitate learning.

## 📚 Table of Contents

1. [Introduction & Variables](#00-introduction--variables)
2. [Arrays & Slices](#01-arrays--slices)
3. [Standard Library](#02-standard-library)
4. [Loops](#03-loops)
5. [Booleans](#04-booleans)
6. [Expressions & Control Flow](#05-expressions--control-flow)
7. [Functions](#06-functions)
8. [Multiple Return Values](#07-multiple-return-values)
9. [Package Scope](#08-package-scope)
10. [Maps](#09-maps)
11. [Pass By Value](#10-pass-by-value)
12. [Pointers](#11-pointers)
13. [Structs & Custom Types](#12-structs--custom-types)
14. [Receiver Functions (Methods)](#13-receiver-functions-methods)
15. [User Input](#14-user-input)

---

## 00. Introduction & Variables

**File:** `00-intro/Intro.go`

### Topics Covered:
- ✅ Variable declaration methods (explicit, implicit, short)
- ✅ Basic data types: `string`, `int`, `float32`, `float64`
- ✅ String concatenation and formatting
- ✅ Printf format specifiers (`%v`, `%q`, `%T`, `%f`)
- ✅ Type information and numeric ranges

### Key Concepts:
```go
// Three ways to declare variables
var name string = "John"      // Explicit
var age = 25                   // Implicit (type inference)
score := 95.5                  // Short declaration (most common)
```

### Format Specifiers:
- `%v` - Default value format
- `%q` - Quoted string
- `%T` - Variable type
- `%f` - Floating-point number
- `%0.2f` - Float with 2 decimal places

---

## 01. Arrays & Slices

**File:** `01-Array/Arrays.go`

### Topics Covered:
- ✅ Fixed-size arrays
- ✅ Dynamic slices
- ✅ Appending to slices
- ✅ Slice ranges and subslicing
- ✅ Array vs Slice comparison

### Key Differences:

| Feature | Array | Slice |
|---------|-------|-------|
| Size | Fixed | Dynamic |
| Declaration | `[3]int{1,2,3}` | `[]int{1,2,3}` |
| Can grow | ❌ No | ✅ Yes (with `append()`) |
| Use case | Rare | Common |

### Slice Operations:
```go
scores := []int{100, 50, 20}
scores = append(scores, 120)        // Add element
middle := scores[1:]                // From index 1 to end
first := scores[:2]                 // First 2 elements
copy := scores[:]                   // Copy entire slice
```

---

## 02. Standard Library

**File:** `02-STD/STD.go`

### Topics Covered:
- ✅ Sorting integers (`sort.Ints()`)
- ✅ Sorting strings (`sort.Strings()`)
- ✅ Searching in sorted slices
- ✅ String manipulation functions (commented examples)

### Sort Package Functions:
```go
ages := []int{45, 20, 30, 56}
sort.Ints(ages)                     // In-place sorting
index := sort.SearchInts(ages, 30)  // Binary search

names := []string{"Bob", "Alice", "Charlie"}
sort.Strings(names)                 // Alphabetical sort
```

### Important Notes:
- ⚠️ `sort.Ints()` modifies the original slice
- ⚠️ `sort.SearchInts()` only works on **sorted** slices
- 💡 Returns insertion index if element not found

---

## 03. Loops

**File:** `03-Loops/Loops.go`

### Topics Covered:
- ✅ Traditional for loop (C-style)
- ✅ Index-based iteration
- ✅ Range-based iteration
- ✅ Underscore `_` for unused variables

### Loop Patterns:

```go
// Traditional for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Range over slice
for index, value := range names {
    fmt.Printf("%d: %s\n", index, value)
}

// Ignore index with _
for _, value := range names {
    fmt.Println(value)
}
```

---

## 04. Booleans

**File:** `04-Bool/bool.go`

### Topics Covered:
- ✅ Boolean type declaration
- ✅ Comparison operators (`==`, `!=`, `<`, `>`, `<=`, `>=`)
- ✅ Boolean expressions
- ✅ Negation with `!` operator

---

## 05. Expressions & Control Flow

**File:** `05-expression/expression.go`

### Topics Covered:
- ✅ If-else statements
- ✅ Logical AND (`&&`) and OR (`||`) operators
- ✅ Loop control (`continue`)
- ✅ Combining conditions

---

## 06. Functions

**File:** `06-Functions/functions.go`

### Topics Covered:
- ✅ Function definitions
- ✅ Parameters and return types
- ✅ Higher-order functions (functions as parameters)
- ✅ Math package usage

```go
// Basic function
func sayHello(name string) {
    fmt.Println("Hello", name)
}

// Function with return value
func circleArea(radius float64) float64 {
    return math.Pi * math.Pow(radius, 2)
}
```

---

## 07. Multiple Return Values

**File:** `07-MultipleReturn/multiplereturn.go`

### Topics Covered:
- ✅ Functions returning multiple values
- ✅ Unpacking multiple returns
- ✅ Using underscore to ignore returns

```go
func getInitials(name string) (string, string) {
    parts := strings.Split(name, " ")
    return parts[0][:1], parts[1][:1]
}

first, last := getInitials("John Doe")
```

---

## 08. Package Scope

**Files:** `08-PackageScope/main.go`, `package1/Package1.go`, `package2/Package2.go`

### Topics Covered:
- ✅ Creating and importing packages
- ✅ Public vs private (exported vs unexported)
- ✅ Package organization

### Key Concepts:
```go
// Uppercase = Public (exported) - accessible from other packages
func PublicFunction() { }
var PublicVar = "visible"

// Lowercase = Private (unexported) - only in this package
func privateFunction() { }
var privateVar = "hidden"
```

### Package Structure:
```
08-PackageScope/
├── main.go           // imports package1 and package2
├── package1/
│   └── Package1.go   // package package1
└── package2/
    └── Package2.go   // package package2
```

---

## 09. Maps

**File:** `09-Maps/maps.go`

### Topics Covered:
- ✅ Creating maps with different key/value types
- ✅ Accessing values by key
- ✅ Updating map values
- ✅ Looping through maps

### Key Concepts:
```go
// Create a map: map[keyType]valueType
menu := map[string]float64{
    "soup": 4.18,
    "rice": 1.98,
}

// Access a value
fmt.Println(menu["rice"])  // 1.98

// Update a value
menu["rice"] = 2.50

// Loop through map
for key, value := range menu {
    fmt.Println(key, value)
}
```

---

## 10. Pass By Value

**File:** `10-PassByValue/PassByValue.go`

### Topics Covered:
- ✅ Go passes copies by default (pass by value)
- ✅ Using pointers to modify original values
- ✅ Returning new values as alternative
- ✅ Maps are reference types (no pointer needed)

### Key Concepts:
```go
// Pass by value - does NOT modify original
func updateName(name string) {
    name = "new"  // Only changes local copy
}

// Pass by pointer - DOES modify original
func updateNamePtr(name *string) {
    *name = "new"  // Changes original
}

// Maps are reference types - no pointer needed
func updateMenu(m map[string]float64) {
    m["coffee"] = 2.50  // Modifies original map
}
```

---

## 11. Pointers

**File:** `11-Pointers/pointer.go`

### Topics Covered:
- ✅ Getting memory address with `&`
- ✅ Dereferencing with `*`
- ✅ Pointer types (`*string`, `*int`, etc.)
- ✅ Modifying values through pointers

### Quick Reference:
```go
name := "mahmoud"

// & = get memory address
fmt.Println(&name)  // 0xc0000140a0

// *type = pointer type
func update(name *string) {
    *name = "new"  // * = value at address
}

update(&name)  // Pass address
```

| Symbol | Meaning | Example |
|--------|---------|---------|
| `&var` | Get address of var | `&name` |
| `*ptr` | Get value at address | `*name = "new"` |
| `*type` | Pointer type | `func(n *string)` |

---

## 12. Structs & Custom Types

**Files:** `12-StructsCustomType/main.go`, `structs/Person.go`

### Topics Covered:
- ✅ Defining structs
- ✅ Creating struct instances
- ✅ Accessing struct fields
- ✅ Constructor functions

### Key Concepts:
```go
// Define a struct
type Person struct {
    Name string
    Age  int
}

// Constructor function (convention: NewTypeName)
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}

// Usage
person := NewPerson("Mahmoud", 20)
fmt.Println(person.Name)  // Access field
```

---

## 13. Receiver Functions (Methods)

**Files:** `13-ReciverFunctionsWithAndWithoutPointers/main.go`, `structs/Person.go`

### Topics Covered:
- ✅ Value receivers (work on copy)
- ✅ Pointer receivers (modify original)
- ✅ Public vs private methods
- ✅ Method chaining patterns

### Key Concepts:
```go
// Value receiver - works on a COPY
func (p Person) GetName() string {
    return p.Name
}

// Pointer receiver - modifies ORIGINAL
func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}

// Usage
person := NewPerson("Mahmoud", 20)
person.UpdateAge(25)  // Go auto-converts to pointer
```

### When to Use:
| Receiver Type | Use When |
|---------------|----------|
| Value `(p Person)` | Read-only, small structs |
| Pointer `(p *Person)` | Need to modify, large structs |

---

## 14. User Input

**Files:** `14-UserInput/main.go`, `structs/Person.go`

### Topics Covered:
- ✅ Reading from stdin with `bufio`
- ✅ Parsing user input
- ✅ Input validation loops
- ✅ Converting strings to other types

### Key Concepts:
```go
// Create a reader for keyboard input
reader := bufio.NewReader(os.Stdin)

// Read until Enter is pressed
input, _ := reader.ReadString('\n')

// Clean up the input
input = strings.TrimSpace(input)

// Convert to other types
age, _ := strconv.Atoi(input)  // string to int
```

### Common Functions:
| Function | Purpose |
|----------|---------|
| `bufio.NewReader(os.Stdin)` | Create input reader |
| `reader.ReadString('\n')` | Read until Enter |
| `strings.TrimSpace(s)` | Remove whitespace |
| `strings.ToUpper(s)` | Convert to uppercase |
| `strconv.Atoi(s)` | String to int |

---

## 🚀 Getting Started

### Prerequisites
- Go 1.25 or higher installed
- Basic understanding of programming concepts

### Running the Examples

```bash
# Navigate to a module folder
cd 00-intro

# Run the Go file
go run Intro.go
```

---

## 📖 Learning Path

**Recommended Order:**

1. **00-intro** → Variables and types
2. **01-Array** → Collections
3. **03-Loops** → Iteration
4. **04-Bool** & **05-expression** → Logic
5. **06-Functions** & **07-MultipleReturn** → Code organization
6. **08-PackageScope** → Project structure
7. **09-Maps** → Key-value data
8. **10-PassByValue** & **11-Pointers** → Memory concepts
9. **12-Structs** & **13-ReceiverFunctions** → Custom types
10. **14-UserInput** → Interactive programs

---

## 💡 Key Takeaways

| Concept | Summary |
|---------|---------|
| Variables | Use `:=` for new, `=` for existing |
| Collections | Prefer slices over arrays |
| Loops | Only `for` exists, use `range` |
| Functions | Can return multiple values |
| Packages | Uppercase = public, lowercase = private |
| Maps | Reference type, key-value pairs |
| Pointers | `&` = address, `*` = value at address |
| Structs | Group related data, use constructors |
| Methods | Pointer receiver to modify original |

---

*Last Updated: January 13, 2026*