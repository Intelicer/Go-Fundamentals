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
- ✅ While-style loop (condition only)
- ✅ Index-based iteration
- ✅ Range-based iteration
- ✅ Underscore `_` for unused variables

### Loop Patterns:

#### Traditional For Loop
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

#### While-Style Loop
```go
x := 1
for x <= 5 {  // No "while" keyword in Go
    fmt.Println(x)
    x++
}
```

#### Range Over Slice
```go
names := []string{"Alice", "Bob", "Charlie"}

// With index and value
for index, value := range names {
    fmt.Printf("%d: %s\n", index, value)
}

// Value only (ignore index)
for _, value := range names {
    fmt.Println(value)
}

// Index only (ignore value)
for index := range names {
    fmt.Println(index)
}
```

---

## 04. Booleans

**File:** `04-Bool/bool.go`

### Topics Covered:
- ✅ Boolean type declaration
- ✅ Comparison operators
- ✅ Boolean expressions
- ✅ Negation with `!` operator

### Comparison Operators:

| Operator | Meaning | Example |
|----------|---------|---------|
| `==` | Equal to | `age == 23` |
| `!=` | Not equal | `age != 23` |
| `<` | Less than | `age < 30` |
| `>` | Greater than | `age > 20` |
| `<=` | Less or equal | `age <= 50` |
| `>=` | Greater or equal | `age >= 18` |

### Boolean Logic:
```go
isAdult := age >= 18
hasLicense := true

if isAdult && hasLicense {
    fmt.Println("Can drive")
}

if !isAdult {  // Negation
    fmt.Println("Not an adult")
}
```

---

## 05. Expressions & Control Flow

**File:** `05-expression/expression.go`

### Topics Covered:
- ✅ If-else statements
- ✅ Logical AND (`&&`) operator
- ✅ Logical OR (`||`) operator
- ✅ Loop control (`continue`, `break`)
- ✅ Combining conditions

### Logical Operators:

#### AND (`&&`) - Both must be true
```go
if age > 20 && age < 30 {
    fmt.Println("Age is 20-30")
}
```

**Truth Table:**
| A | B | A && B |
|---|---|--------|
| T | T | **T** |
| T | F | F |
| F | T | F |
| F | F | F |

#### OR (`||`) - At least one must be true
```go
if isWeekend || isHoliday {
    fmt.Println("Can rest!")
}
```

**Truth Table:**
| A | B | A \|\| B |
|---|---|----------|
| T | T | **T** |
| T | F | **T** |
| F | T | **T** |
| F | F | F |

### Loop Control:
```go
// continue - skip to next iteration
for i, val := range names {
    if i == 1 {
        continue  // Skip this iteration
    }
    fmt.Println(val)
}

// break - exit loop entirely
for i, val := range names {
    if i == 1 {
        break  // Stop loop
    }
    fmt.Println(val)
}
```

---

## 06. Functions

**File:** `06-Functions/functions.go`

### Topics Covered:
- ✅ Function definitions
- ✅ Parameters and return types
- ✅ Functions with return values
- ✅ Higher-order functions (functions as parameters)
- ✅ Math package usage

### Function Syntax:
```go
// Basic function (no return)
func sayHello(name string) {
    fmt.Println("Hello", name)
}

// Function with return value
func circleArea(radius float64) float64 {
    return math.Pi * math.Pow(radius, 2)
}

// Higher-order function (takes function as parameter)
func process(items []string, fn func(string)) {
    for _, item := range items {
        fn(item)  // Call the passed function
    }
}
```

### Usage Examples:
```go
// Call basic function
sayHello("Alice")

// Store returned value
area := circleArea(10.5)

// Pass function as argument
process(names, sayHello)  // No () after sayHello
```

---

## 07. Multiple Return Values

**File:** `07-MultipleReturn/multiplereturn.go`

### Topics Covered:
- ✅ Functions returning multiple values
- ✅ Unpacking multiple returns
- ✅ Using underscore to ignore returns
- ✅ String slicing and manipulation
- ✅ Handling edge cases

### Multiple Return Syntax:
```go
// Function returning TWO values
func getInitials(name string) (string, string) {
    parts := strings.Split(name, " ")
    if len(parts) > 1 {
        return parts[0][:1], parts[1][:1]
    }
    return parts[0][:1], "_"
}

// Unpack both values
first, last := getInitials("John Doe")  // first="J", last="D"

// Ignore one value
first, _ := getInitials("Madonna")      // first="M", ignore second
```

### Common Pattern - Error Handling:
```go
// Common Go idiom: return value and error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Usage
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

---

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher installed
- Basic understanding of programming concepts

### Running the Examples

1. **Navigate to a module folder:**
   ```bash
   cd 00-intro
   ```

2. **Run the Go file:**
   ```bash
   go run Intro.go
   ```

3. **Experiment:**
   - Uncomment different functions in `main()`
   - Modify values and observe outputs
   - Read the inline comments for explanations

---

## 📖 Learning Path

**Recommended Order:**

1. Start with **00-intro** to understand variables and types
2. Move to **01-Array** for collections
3. Learn **03-Loops** for iteration
4. Study **04-Bool** and **05-expression** for logic
5. Master **06-Functions** and **07-MultipleReturn** for code organization
6. Explore **02-STD** for standard library usage

---

## 💡 Key Takeaways

### Variable Declaration
- Use `:=` for new variables (most common)
- Use `=` for existing variables
- `var` keyword for explicit types or zero values

### Collections
- **Arrays**: Fixed size, rarely used
- **Slices**: Dynamic, use `append()` to grow

### Loops
- Only `for` loop exists in Go (no `while`)
- `range` is preferred for iterating collections
- Use `_` to ignore unused values

### Functions
- Can return multiple values (unique to Go)
- Functions are first-class (can be passed as parameters)
- Named return values are possible but optional

### Best Practices
- ✅ Prefer short variable declaration `:=`
- ✅ Use `range` for cleaner loops
- ✅ Handle multiple returns explicitly
- ✅ Use meaningful variable names
- ✅ Comment complex logic

---

## 📚 Additional Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Playground](https://go.dev/play/) - Test code online

---

## 🤝 Contributing

Feel free to:
- Add more examples
- Improve documentation
- Fix errors or typos
- Suggest new topics

---

## 📝 Notes

- All code is heavily commented for educational purposes
- Examples progress from simple to complex
- Each file is self-contained and runnable
- Code follows Go conventions and best practices

---

**Happy Learning! 🎉**

*Last Updated: January 12, 2026*