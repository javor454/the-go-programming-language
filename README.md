# Notes
- Control flow - determines the execution order of statements (if, for...)
- To write efficient code, you need to be aware of the lifecycle of variables
- if file contains `init()` func, its automatically executed on program start

## Variable
- each variable has a type and a value


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
line := "ahoj"
m[line]++ // ["ahoj":1] 

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

## Pointers
- address of a location, where the data of a variable is stored in memory
- i can get the address of variable:
    - using ampersand &
    - by creating unnamed variable using new keyword, value is initialized to zero value of type `a = new(int) // *a = 0`
- i can get the value of pointer with star `fmt.Println(*a) // 0`

## Integers
- unsigned integers used mainly for bitmasks
- octal numbers used mainly for file permissions

## Runes
- alias for `int32`
- represents single unicode code point which maps to specific character
- each rune can be represented by 1-4 bytes

## Strings
- immutable sequence of bytes `[]byte`
```go
    s := "abc"
	s[0] = "L" // compile error: cannot assign to s[0] 
```
- UTF8 encoded sequences of runes
- `len(str)` returns number of bytes, not characters
  - `fmt.Println(len("łż")) // 4` - each char has 2 bytes
  - `fmt.Println(len("世界")) // 6` - each char has 3 bytes
- `utf8.RuneCountInString("łż") // 2` returns number of runes in string
- `len([]rune("łż")) // 2` returns number of elements in slice
```go
    for i, _ := range "łż" {
		fmt.Println(i) // 0, 2
	}
```
```go
    s := "12345"
	fmt.Println(s[1:4]) // "234"
	fmt.Println(s[:3]) // "123"
	fmt.Println(s[3:]) // "45"
````
- string enclosed with "" brackets interprets escape handles, `` ignores them
- `` may spread over several lines or show backslashes \ because it represents them literally
- important string pkgs: strings, bytes, strconv, unicode
- because strings are immutable, building up strings incrementally can involve a lot of allocation and copying
  - its more efficient to use `bytes.Buffer`
- string x []byte conversion
  - converting string to []byte allocates new byte array holding a copy of the bytes of s and yielding a sice that references that whole array
  - converting []byte to string also makes a copy
```go
	numbers := []int{1, 2, 3, 4, 5}
	var buf bytes.Buffer
	buf.WriteByte('[') // or WriteRune
	for i, v := range numbers {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	fmt.Println(buf.String()) // "[1, 2, 3, 4, 5]"
```
- string x number conversion
  - strconv.Itoa (int to ascii)
  - strconv.FormatInt (convert to different base)
  - fmt.Sprintf
- number x string parsing
  - strconv.Atoi
  - strconv.ParseInt

## Constants
- value known at compile time
- can be boolean, string, int
- if type is missing is infered from the expression
- untyped constants have bigger precision e.g. untyped int has 256 bits

### iota constant generator
- creates sequence of related values

## Arrays
- has fixed length, must be known at compile time
- is passed by value
- has one type
- len() returns length
- if values are not initialized, they are equal zero value
  - `var a [3]int // [0,0,0]`
  - `var a [3]int = [3]int{1,2,3}// [1,2,3]`
- if ... is used instead of length, length of array is determined from initializers
  - `a := [...]int{1,2,3} // len(a) = 3`
- if value with index 5 is added, values until that key is filled up to array
  - `a := [...]int{5: 1} // [0 0 0 0 0 1]`
- if array type is comparable then the array is comparable too, can be compared with == to check if it has same elements
```
    a := [...]int{4: 1}
	b := [5]int{4: 1}
	fmt.Printf("%v", a == b) // true
```
- slice operator - if slicing `original := make([]int, 0, 10)`
  - `arr[1:]` - get all elements from 1st to the end of array, has capacity 9: 10-1
  - `arr[:9]` - get all elements from begin to 9th element, has capacity 10: no offset from start 
  - `arr[3:5]` - get all elements from 3rd to 5th element, has capacity 7: 10-3
  - NOTICE the capacity, its copied from original slice but shortened from beginning¨
  - slicing beyond CAP causes PANIC
  - slicing beyong LEN extends the slice

## Slices
- has variable length
- has header which consists of pointer, length and capacity
  - pointer - points to underlying array
  - length - number of elements in array
  - capacity - reserved memory for all elements
- when passing slices as argument, the header is passed as value but the reference to the array stays
- are not comparable by ==, comparison must be done manually
- zero value of slice is nil
- nil slice
  - has no underlying array
  - has zero length and capacity
- check if slice nil `sliceA == nil`
- check if slice empty `len(sliceA) == nil`
- `copy` - copies values from one slice to another of same type
```go
// how to pass slice as reference, if not then the header is passed as value even tho it contains reference to underlying array
func main() {
	a := []int{1, 2, 3, 4, 5}
	remove(&a, 2)
	fmt.Printf("%v", a)
}

func remove(slice *[]int, i int) {
	copy((*slice)[i:], (*slice)[i+1:])
	(*slice) = (*slice)[:len((*slice))-1]
}
```

## Maps
- hash table
- unordered collection of key / value pairs where keys are unique
- get, set, update can be done in constant time
- zero value is nil
- when checking if value exists, use this notation, otherwise zero value is returned
```go
    a := make(map[int]int)
	if val, ok := a[1]; ok {
		fmt.Printf("%v\n", val)
	} else {
		fmt.Println("nok")
	}
```
- struct which contain slices, maps, or functions (comparable struct types) CANNOT be used as key in map, otherwise can be

## Struct
- aggregated data type that groups together 0..N fields with different types
- field is a named value
- passed by value
- create by:
  - `a := Krysa{}`
  - `a := new(Krysa)`
- struct cannot contain field with same type as itself, only a reference to the same type as itself
- empty struct `struct{}`
```go
    type Krysa struct {
		vek int
	}

	krysak := Krysa{}
	fmt.Printf("%v\n", krysak.vek) // 0
```
- is comparable by ==
- embedding structs as fields
```go
type Point struct {
	X, Y float64
}
type Circle struct {
	Point
	R float64
}

var c = Circle{
	Point: Point{X: 0, Y: 0},
	R:     0,
}

fmt.Println(c.X, c.Y, c.R)
```

### JSON
- standard notation for sending / receiving structured information
- uses reflection
- convert go structures to JSON via marshalling - produces byte slice
  - use field tags in structs - metadata associated to fields in compile time
```go
type Str struct {
	Point int `json:"point"`
}
a = Str{Point:5}
data, err := json.Marshall(a)
if err != nil {
	return error.New("marshall failed")
}
```
  - only exported fields are marshalled
  - use MarshalIndent for human readable json
- convert JSON to go structure via unmarshalling
```go
type Str struct {
	Point int `json:"point"`
}
var objs []Str
if err := json.Unmarshall(data, &objs); err != nil {
	
}
```
- Encoder / Decoder - process json stream by chunks

### Text and HTML templates
- string or file which contains actions enclosed in brackets {{  }}
  - actions trigger behaviour such as:
    - printing values
    - selecting struct fields
    - if else
    - loops
    - calling fns and methods and using other templates
- possible to pipe output of one operation to another
```go
// template
{{ .Title | printf "%.64s" }} // equal to fmt.Sprintf
{{ .CreatedAt | daysAgo }} // uses function daysAgo

...

func daysAgo(t time.Time) int {
  return int()time.Since(t).Hours() / 24)
}
```
- templates workflow 2-step process:
  1. parse template into internal representation
  2. execute on specific inputs
```go
report, err := template.New("report").
    Funcs(template.FuncMap{"daysAgo", daysAgo}).
	Parse(templ)
if err != nil {...}
```
#### HTML templates
- escaping for additional injection attack protection
- has safe custom types
  - template.HTML for trusted HTML
  - string for untrusted plain text

## Functions
- wraps sequence of statements as a unit which can be called from elsewhere in a program multiple times
- arguments are passed by value, so the func receives a copy
  - except reference like pointer, slice, map, function or channel
- go doesnt have default values
- functions in source without body are implemented in different language (e.g. assembly)
- function can return one or more named or unnamed return values
- zero value of function type is nil
- calling a nil function causes panic
- named functions can be declared only at package level
- functions are not comparable
- functions are considered a reference type

### Recursion
- function which calls itself
- GOlang stack has variable size and grows up to 1GB of size

### Errors
- last parameter in function
- if error is not nil, other returned parameters might be usually ignored if theres not a requirement to return partial data (bytes written etc. )
- because errors are usually chained, messages should not be capitalized and should avoid newlines
- error messages should be consistent
- in case the problem is transient (přechodný) it might make sense to retry (backoff strategy etc.)
- in case its impossible to continue (bug), the caller can log error and stop the program gracefully although its best practice to return the information to the caller OR continue with limited functionality
- if the problem is one cause use `ok`, in case there might be several causes to the issue use `err`
```go
value, ok := cache.Lookup(key)
if !ok {...}
resp, err := http.Get(url)
if err != nil {...}
```
- error je interface

### Anonymous functions (closure) 
- fixed in 1.22
- has access to entire lexical environment (can access variables outside of its scope - these can be garbage collected as soon as all the references to the returned functions are gone)
- captures values outside of its lexical scope by reference
```go
func createCounter() func() int {
    count := 0          // count can't be garbage collected
    return func() int { // while this function exists
        count++
        return count
    }
}
```

### Variadic functions
- has variadic number of arguments e.g. `fmt.Printf()`
```go
func sum(vals ...int) int {
	var total int
	for _, val := range vals {
        total += val
    }
	return total
}
values := []int{1,2,3,4}
sum(values...)
```

### Deferred function calls
- call deffered until the function containing the statement finishes:
  - succesfully on return
  - panic 
- any number of calls can be deffered, they will be executed in the reverse order they were called
- usually used to ensure release of resources right after acquire of resources upon operations like:
  - open/close
  - connect/disconnect
  - lock/unlock

### Panic
- go runtime panics in case it runs into problems at runtime like out-of-bounds array access or nil pointer dereference
- usually functions with name prefix Must panic in case of error
- when panic occurs, all deffered funcs run in reverse order, program is terminated and the stack is printed on standard error output

#### Recover
- panics can be recovered using the recover func in defer of the func which panicked
- this will stop the panic and return its value, then the func that panicked returns normally instead of continuing where it left off

## Methods
- function associated with particular type
- method can be created on any named type which is not pointer nor interface
```go
type IntSlice []int

// Sum calculates the sum of all elements in the slice
func (s IntSlice) Sum() int {
    total := 0
    for _, v := range s {
        total += v
    }
    return total
}
```
- pointer receiver is used when method needs to update the object its associated to
- if one method has pointer receiver, then all other methods should have it too
- receiver can be nil in case of pointer receivers, value receivers always need value
- encapsulation
  - using private fields and methods in combination with public methods as getters and setters
- composition
  - using embedding structs as fields e.g.
```go
type Point struct {
	X, Y float64
}
type Circle struct {
	Point // or *Point
	R float64
}
```
  - embedded structs can be pointers
  - method can be used as a value e.g.
```go
p := Point{1, 2}
met := p.ScaleBy
met(2)
```
- when naming getters, we usually omit the word "Get"
  - other prefixes are also ommited like Fetch, Find, Lookup

## Interfaces
- set of methods which describe behavior of a type - it is enough that type implements all methods of interface
- satisfied implicitly
- can be implemented by any type
- interface type is a set of methods that a concrete type must implement
- supports embedding (composition)
- type satisfies interface if it implements all methods of interface
- check if type satisfies interface `var _ io.Writer = (*bytes.Buffer)(nil)`
- zero value of interface is nil (nil value and type)

### Error interface
- error is an interface with method `Error() string`
- when created by New, underlying type is `*errors.errorString` = error created with same value is always different
  - fmt.Println(errors.New("EOF") == errors.New("EOF")) // false
- more popular is fmt.Errorf which formats error message and returns it

### Type assertion
- checks if the value is of a certain type
- if not, it panics
- if yes:
  - if the asserted type is a concrete type - extracts the value to a new variable of asserted type
  - if the asserted type is an interface - makes more methods available, but the value is still of original type
  - if the asserted type is a nil interface - it panics
- to prevent panic, use comma ok idiom
```go
var w io.Writer
w, ok := r.(io.Writer)
if !ok {
    fmt.Println("r does not satisfy io.Writer")
}
```
- is often used in switch statements
```go
switch x.(type) {
case nil:
    // ...
}
```

## Goroutines and Channels
- communicating sequential processes (CSP)
  - model of concurrency where values are passed between independent activities (goroutines) but variables are mostly confined to single activity
- shared memory multithreading

### Goroutines
- concurrently executing activity
- main function is also a goroutine
- new goroutine is created with `go` keyword
- when main function returns, all goroutines are killed

### Channels
- connections between activities
- communication mechanism which lets goroutines communicate with each other
- channels are typed, can be used to send and receive values with the same type
- created with `make` keyword
```go
ch := make(chan int)
```
- reference type
- send statement
```go
ch <- 1
```
- receive statement
```go
x := <-ch
<-ch // receive, ignore value
```
- close = closes channel, no more sends allowed, if send then panic
- receive from closed channel yields the values already send until no more values are left
- channels can be closed with `close` keyword
```go
close(ch)
```

#### Unbuffered channels
- created with `make(chan int)`
- send operation blocks the sending goroutine until the receive operation is performed by another goroutine
- receive operation blocks the receiving goroutine until the send operation is performed by another goroutine
- use struct{}{} as a value to send to channel to signal completion
- use ok idiom to check if channel is closed
```go
x, ok := <-ch
if !ok {
    fmt.Println("channel is closed")
}
```
- when ranging over channel, the loop will exit when the channel is closed
```go
for x := range ch {
    fmt.Println(x)
}
```

#### Pipelines
- connected set of stages (goroutines) by channels

#### Unidirectional channels
- send only `chan<-`
- receive only `<-chan`
- violation of this rule will cause compile error
- close on send only channel will cause compile error
- can convert bidirectional channel to unidirectional, but not the other way around

#### Buffered channels
- has queue of elemements, determined by capacity `make(chan int, 100)`
- send operation inserts element at the end of queue
- receive operation removes element from the beginning of queue
- if the buffer is full, the sending operation blocks until the buffer has space
- if the buffer is empty, the receiving operation blocks
- goroutine leak
  - if the buffer is too small, the sending goroutine will block and the goroutine will leak - will not be automatically collected by GC
- waitgroup
  - used to wait for a set of goroutines to finish
  - created with `sync.WaitGroup`
  - add number of goroutines to wait for with `Add`
  - done with `Done`
  - wait with `Wait`

#### Multiplexing using select
- used to wait on multiple channels
- select statement
```go
select {
case <-ch1:
    // ...
case <-ch2:
    // ...
default:
    // ...
}
```
- default case is optional
- if default case is present, it will be executed if no other case is ready
- if default case is not present, the select will block until one of the cases is ready
- if multiple cases are ready, one is chosen at random
- select{} waits forever
- channel zero value is nil
  - nil channel will block forever

## Concurrency with shared variables

### Race condition
- data race - when two goroutines access the same variable concurrently and at least one of them is a write
- do not communicate by sharing memory, share memory by communicating
- monitor goroutine
  - goroutine which brokers access to variable using channel requests
- detect race condition with `go run -race`

### Mutex
- mutual exclusion = lock
- used to prevent race condition
- lock and unlock should be used in pairs, or with defer
- if lock is acquired by goroutine, the next goroutine which tries to acquire gets blocked until the lock is released

### Deadlock
- A deadlock is a situation in concurrent programming where two or more processes are unable to proceed because each is waiting for resources held by another process, creating a circular dependency that prevents any progress.

#### RWMutex
- read-write Mutex
- allows multiple readers or a single writer

### Sync.once
- used to execute a function only once
- implemented by mutex

## Goroutines and threads
- OS threads
  - fixed size block of memory 2MB for its stack
- goroutines
  - dynamic size starts at 2KB
  - grow and shrink as needed
  - have no identity

### Goroutine scheduling
- OS threads are scheduled by OS kernel
- go runtime multiplexes (schedules) N goroutines on M OS threads

### GOMAXPROCS
- controls how many OS threads are used by the Go runtime to run goroutines
- default is number of available CPU cores
- can be set with `runtime.GOMAXPROCS(n)`

## Packages and the go tool
- brings modularity by packaging reusable code together
- package is a collection of source files in the same directory
- each directory contains at least one file with package clause `package main`
- package name is the directory name
- package declaration must be the first line in the source file
- package can be imported using `import` keyword
- import can be used to import packages

## Testing
- package `testing`
- `go test` command
- `go test -bench` command
- `go test -cover` command
- `go test -v` command
- `go test -run` command

### Test files
- must be in the same directory as the file it is testing
- must be named `*.test.go`
- must import `testing` package
- must contain `Test` prefix
- must be in the same package as the file it is testing


## Low level programming
- unsafe package
  - allows to bypass type safety of Go
  - provides low level functions for manipulating pointers
  - can cause program to crash
  - can cause data races
  - can cause memory leaks
- calling c code from go




## Interesting packages
- url.QueryEscape - encode special characteres for safe use in URL
- pkg path for manipulation with URLS
  - path/filepath for manipulation with filenames
## Performance tips
- keeping pointers to short-lived objects inside long-lived objects (global vars) will prevent GC to reclaim the short-lived objects
- because strings are immutable, building up strings incrementally can involve a lot of allocation and copying
  - its more efficient to use `bytes.Buffer`
- string x []byte conversion
  - converting string to []byte allocates new byte array holding a copy of the bytes of s and yielding a sice that references that whole array
  - converting []byte to string also makes a copy
  - pouzit unsafe a pretypovat

## Best practices
- successful execution path of code should not be indented
- because errors are usually chained, messages should not be capitalized and should avoid newlines
- error should be the last parameter in function return params
- in case error is returned by function, other returned parameters can be ignored if the requirements doesn't state otherwise
- in case the problem is transient (přechodný) it might make sense to retry (backoff strategy etc.)
- in case its impossible to continue in execution (bug), the caller can log error and stop the program gracefully although its best practice to return the information to the caller OR continue with limited functionality
- if one method has pointer receiver, then all other methods should have it too
- when naming getters, we usually omit the word "Get"
- ticker best practice against goroutine leak
```go
ticker := time.NewTicker(time.Second)
<-ticker.c // receive from ticker channel
ticker.Stop() // cause the ticker goroutine to stop
```

## Interesting problems

### Iteration variable capture
- This happens because Go's for loop reuses the same variable for each iteration. Closures capture the variable reference, not its value at creation time
```go
// fc will create and later remove all dirs

// bad
var rmdirs []func()
for i := 0; i<len(dirs); i++ {
	os.MkdirAll(dirs[i], 0755) // ok
	rmdirs = append(rmdirs, func() {
	    os.RemoveAll(dirs[i]) // incorrect, after loop is done, i will hold final loop i value	
    })
}

//good
var rmdirs []func()
for i := 0; i<len(dirs); i++ {
	helpVar := i // helper var will capture value of i even after loop end
    os.MkdirAll(dirs[i], 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dirs[helpVar])	
    })
}
```

### Tracing using defer and closure
```go
func bigSlowOperation() {
    defer trace("bigSlowOperation")()
    time.Sleep(3 * time.Second)
}

func trace(msg string) func() {
    start := time.Now()
    fmt.Printf("enter %s\n", msg)
    return func() { fmt.Printf("exit %s (%s)\n", msg, time.Since(start)) }
}
```
