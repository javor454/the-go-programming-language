# Notes
- Control flow - determines the execution order of statements (if, for...)
- To write efficient code, you need to be aware of the lifecycle of variables
- if file contains `init()` func, its automatically executed on program start



## Map
- map is a reference type - points to the underlying data structure
    - if i pass a map to a function, the function can modify the map
- set of key-value pairs
- provides constant time operations for adding, removing, and finding elements
- key can be of any type that can be compared with ==
- if i print the map, the order of elements is random
### Examples
```go
// add element to map using shorthand
m := make(map[string]int)
line := "koko"
m[line]++
// ["ahoj":1] 

// iterate through map with range loop
for line, n := range m {
    fmt.Printf("%s: %d\n", line, n)
}
```

## Package
- functions and package level entities are exported if their name starts with a capital letter
- functions and package level entities may be declared in any order
- if an entity is declared in function - its local to this function
- if an entity is declared outside of function - its visible to all files in the package it belongs
- package level variable lifetime is entire execution of program

## Consts
- values fixed at compile time

## Functions
- execution ends either when calling return or when the function reaches end

## Declaration
- if the initializer part of declaration is omited, the initial value of the variable is set to the zero value of its type
    - nil: interfaces, reference types (slice, map, pointer, channel, func)
    - "": string
    - 0: number
    - false: bool
    - special case -> struct has zero value set for each property based on the property type
- package level variables are initialized before main begins
- local func variables are initialized when the execution reaches them
- scope: region of program text, compile-time property
- lifetime: range of time during program execution when the variable can be refered to by other parts of the program, run-time property

## Pointers
- address of a location, where the data of a variable is stored in memory
- i can get the address of variable:
    - using ampersand &
    - by creating unnamed variable using new keyword, value is initialized to zero value of type `a = new(int) // *a = 0`
- i can get the value of pointer with star `fmt.Println(*a) // 0`

## Integers
- unsigned integers used mainly for bitmasks
- octal numbers used mainly for file permissions

## Assignment
- Tuple assignment
    - assign multiple return values from a single func call OR multiple single value expressions
    - all expressions on the right side are evaluated before assignment to the vars on left side
    - map lookup, type assertion, channel receive produces bool result
        - val, ok = m["key"]
        - val, ok = a.(int)
        - val, ok = <-ch
  ```go
  package main
  import "fmt"
  
  func main() {
    a, b := func() (bool, bool){
        return true, false
    }()
    fmt.Println(a, b)

   	i, j := 1, 2
	i, j = j, i
    fmt.Println(i, j) //i = 2, j = 1
  }
  ```

## Performance tips
- keeping pointers to short-lived objects inside long-lived objects (global vars) will prevent GC to reclaim the short-lived objects

## Best practices
- successful execution path of code should not be indented