```table-of-contents
title:
style: nestedList # TOC style (nestedList|nestedOrderedList|inlineFirstLevel)
minLevel: 0 # Include headings from the specified level
maxLevel: 0 # Include headings up to the specified level
include: 
exclude: 
includeLinks: true # Make headings clickable
hideWhenEmpty: false # Hide TOC if no headings are found
debugInConsole: false # Print debug info in Obsidian console
```

---
# A few points about GoLang [^1]

1. Statically typed language (types can't be changed (not without type conversion))
2. Strongly typed language (can't add int with a string)
3. It's compiled
4. Compilation is fast
5. GoRoutines (builtin concurrency)
6. Simple syntax
7. Garbage Collector

Packages have a bunch of go files
Modules have a bunch of packages

```bash
go mod init <global url : typically github (github.com/abc/abc)>
```

go.mod has the name of the module, go version, and requirements


## Introduction

`./cmd/tutorial/main.go`

We'll have a `package main` line written in the top, signifying the package. Also, `main` is a special package name, that makes the compiler look for the entry point function (main function)

### Functions

> [!info] "fmt" package
> It's a standard library used for formatted input and output operations

```go
import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

```bash
go build cmd/tutorial/main.go
./main

or

go run cmd/tutorial/main.go
```

### Variables

```go
package main

import "fmt"

func main() {
	var intNum int = 32767
	// int, int8, int16, int32, int64, uint, ..., float32, float64
	// string
	
	// we can't perform any arithmetic operation between any 2 types, like int + float32 won't work. I'll have to do this
	//  floatNum32 + float32(intNum32)
	
	var myString string = "Hello World" // single line
	var myString2 string = `Hello
	World`
	
	fmt.Println(len("test")) // 4, but this isn't the number of chars, it's the number of bytes. For ascii, it's 1 byte per char only, but go uses utf8, so, it's slightly variable, for chars outside it
	
	// for actual length, 
	// import "unicode/utf8"
	// fmt.Println(utf8.RuneCountInString("Test"))
	// rune type = character in utf8 
	var myRune rune = 'a'
	fmt.Println(myRune) // 97
	// bool -> true / false
}
```

we don't need to initialize values btw. Go assigns default values too 
For int, int8, ..., uint, uint8,..., float32, ..., rune = 0
For string, it's an empty string, and for booleans, it's false

We can also infer the type, 
```go
var myString = "hello"

// we can also, remove the var keyword

myString := "hello"

// we can also do this

var var1, var2 int = 1, 2
var var3, var4 = 1, 2
var5, var6 := 1, 2
fmt.Println(var1,var2) // 1 2
```

Instead of var, we can use const as well btw. Const = can't be changed. We have to initialize the value explicitly.

---
# Functions & Control Structures

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello"
	printMe(printValue)
	var num int = 11
	var den int = 2
	var result = intDivsion(num, den)
	fmt.Println(result)
	var res, rem, err = intDivisionBoth(num, den)
	if err != nil {
		fmt.Printf(err.Error())
	} else if rem == 0 {
		fmt.Printf("The result of the int division is %v", res)
	} else {
		fmt.Printf("The result of the int division is %v with remainder %v", res, rem)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int) int {
	var result int = num/den
	return result
}

func intDivisionBoth(num int, den int) (int, int, error) {
	var err error // defualt = nil
	if den == 0 {
		err = errors.New("Can't divide by zero")
		return 0, 0, err
	}
	var result int = num/den
	var rem = num%den
	return result, rem, err
}
```

---
# Arrays, Slices, Maps, and Loops

## Arrays

```go
package main
import "fmt"

func main() {
	var intArr [3]int32
	intArr[1] = 123
	fmt.Println(intArr[0]) // 0
	fmt.Println(intArr[1:3]) // [123 0]
	// index 1 and 2
	
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}
```

```go
package main
import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1,2,3}
	fmt.Println(intArr)
}
```

```go
package main
import "fmt"

func main() {
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)
}
```

## Slices

They are wrappers around arrays. They are literally the golang equivalent of vectors in c++.
They are dynamically sized arrays. However, a slice is a reference into an underlying array, while the std::vector object in c++ owns the underlying storage. 

Slice = <pointer, length, capacity>. It's just a descriptor

```go
func main() {
	// by ommitting the length value in the array, we'll get a slice
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice) // [4 5 6]
	
	fmt.Printf("length: %v, capacity: %v\n", len(intSlice), cap(intSlice))
	// 3 3 : [4 5 6]
	
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice) // [4 5 6 7]
	// It checks if the array has enough space for more values. 
	// If not, a new array is created with enough space, and the data is copied
	
	fmt.Printf("length: %v, capacity: %v", len(intSlice), cap(intSlice))
	// 4 6 : [4 5 6 7 * *]
	
	
	
	// we can also append more values
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // intSlice2... : spread operator
	fmt.Println(intSlice)
	
	
	// We can use the make function to also create a slice
	// type, length, (and optionally, capacity)
	var intSlice3 []int32 = make(int32[], 3, 8)
}
```

## Maps

```go
func main() {
	var myMap map[string]uint8 = make(map[string]uint8) 
	// key: string, value: uint8
	
	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"]) // 23
	
	fmt.Println(myMap2["Jason"]) 
	// doesn't exist : we'll get the default value = 0
	
	var age, ok = myMap2["Adam"] // maps return an optional 2nd value (bool)
	// ok = true : the value is in the map, and if it's false, the value isn't
	
	delete(myMap2, "Adam") // if the key doesn't exist, it does nothing : no-op
}
```

## Loops

```go
func main() {
	
	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	for name := range myMap2 {
		fmt.Println(name)
	}
	for name, age := range myMap2 {
		fmt.Println(name, age)
	}
	
	
	intArr := [...]int32{1,2,3}
	for i, v := range intArr { // i = index, v = value
		fmt.Println(i, v)
	}
}
```

```go

// while loops are implemented via the for keyword
i := 0
for i<10 {
	fmt.Println(i)
	i = i + 1
}

// or

i := 0
for {
	if i >= 10 {
		break
	}
	fmt.Println(i)
	i = i + 1
}

// or use a regular for loop

for i:=0; i<10; i++ {
	fmt.Println(i)
}
```

## Strings

If we aren't using plain ascii characters, using a regular string will end up being complicated, as the length will also change. Hence, it's better to use an array of runes, instead of strings in that case. However, runes are just aliases for int32. So, it does end up consuming more memory. 

> [!info]
> Strings are immutable in GoLang

### Runes

> [!info] Printing %T 
> This will print the type of the value

```go
func main() {
	// var myString = "resume"
	var myString = []rune("résumé")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed) // 233, int32
	
	for i, v := range myString {
		fmt.Println(i, v)
	}
	
	var myRune = 'a'
	
	var strSlice = []string{"he", "ll", "o"}
	var catStr = ""
	for i := range strSlice { 
		// range gives two elements : index, and value
		// here we are just using one, so that's the index
		catStr += strSlice[i] 
	}
	fmt.Println(catStr)
	// however, this is inefficient, as we're 
	// literally making a new string in every iteration. We should
	// instead import strings. 
	
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.string()
	fmt.Println(catStr2)
}
```

---
# Structs & Interfaces

## Structs 

Structs = your own type

```go
package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
	// we can also just write "owner" instead of "ownerInfo owner"
	// this will add the subfields directly. So, we'll be able to access
	// name via, gasEngine.name, instead of gasEngine.ownerInfo.name
	
	// we can also just write "int" here, and that will create a subfield named,
	// "int", with the type int. 
	
	// however, if a collision is about to occur, like if name was already a
	// field, we wouldn't be able to do this.
}

type owner struct {
	name string
}

/*
We can also create methods, that are functions directly tied to structs, and can access its values. (e gasEngine) is the part that makes it different from a regular function. Basically the functions within a class
*/
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg:25, gallons:15, ownerInfo:owner{"A"}}
	var myEngine gasEngine2 = gasEngine{25, 15, owner{"A"}}
	gasEngine2.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	
	fmt.Printf("Miles left: %v", myEngine.milesLeft())
	
	var anonymousEngine = struct {
		mpg uint8
		gallons uint8
	}{25,25}
}
```

## Interfaces

Now, we'll also add an electric engine, and use an interface

```go
package main

import "fmt"

type gasEngine struct {
	mpg uint8
	gallons uint8
	ownerInfo owner
	// we can also just write "owner" instead of "ownerInfo owner"
	// this will add the subfields directly. So, we'll be able to access
	// name via, gasEngine.name, instead of gasEngine.ownerInfo.name
	
	// we can also just write "int" here, and that will create a subfield named,
	// "int", with the type int. 
	
	// however, if a collision is about to occur, like if name was already a
	// field, we wouldn't be able to do this.
}

type owner struct {
	name string
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

/*
We can also create methods, that are functions directly tied to structs, and can access its values. (e gasEngine) is the part that makes it different from a regular function. Basically the functions within a class
*/
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}
func (e electricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

/*
Now, this function can take any engine, as long as it has the milesLeft() method in it. Since both gasEngine, and electricEngine have it, it will accept both
*/
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{mpg:25, gallons:15, ownerInfo:owner{"A"}}
	canMakeIt(myEngine, 50)
	var myEngine2 electricEngine = electricEngine{25,15,owner{"A"}}
	canMakeIt(myEngine2, 50)
}
```

---
# Pointer

```go
package main

import "fmt"

func main() {
 var p *int32 = new(int32)
 var i int32
}
```

A lot of the pointer stuff is just like in C. `new()` is just like malloc.

> [!info] Remember about Slices
> Slices are basically just the pointers to the array. So, if u try to make a copy of a slice by, `var sliceCopy = slice`, both sliceCopy and slice will point to the same memory location 

> [!info] %p prints the pointer address

```go
package main

import "fmt"

func main() {
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("%p", &thing1)
	
	var result = square(&thing1)
	fmt.Printf("The result is %v\n", result)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("%p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}
```

---
# Go Routines

> [!info] Imp thing about Concurrency
> Concurrency != Parallelism

Here we use wait groups, (from the sync package). These work just like semaphores, but instead of decrementing the counter, it increases the count when some resource gets added to it. We also have to use the `go` call to make it run concurrently. There's also the Wait() function, that will wait till the values don't become 0. We can't avoid using the wait groups, as without them, the cpu just adds the concurrency, but doesn't actually wait for the tasks to finish.

Without Wait Groups
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1","id2","id3","id4","id5"}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		go dbCall(i)
	}
	fmt.Println(time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println(dbData[i])
}
```

If we run this code, nothing will happen. Our program spawned these tasks in the background, didn't wait for them to finish, and then exited the program. Hence, we need to use wait groups

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1) // increments 1 to the counter
		go dbCall(i)
	}
	wg.Wait() // waits for the counter to go back to 0
	fmt.Println(time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println(dbData[i])
	wg.Done() // decrements 1 from the counter
}
```

Ight, now that we've implemented wait groups, lets move to another important OS concept, "Mutexes"


```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1) // increments 1 to the counter
		go dbCall(i)
	}
	wg.Wait() // waits for the counter to go back to 0
	fmt.Println(time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println(dbData[i])
	m.Lock()
	results = append(results, db) // If used without locks, this can cause a race condition (OS concept)
	m.Unlock()
	wg.Done() // decrements 1 from the counter
}
```

We also have read write mutexes in golang. `sync.RWMutex{}`
In this, we have the Lock and Unlock methods too. But along with them, we also have Read Lock, (`RLock()`), and Read Unlock (`RUnlock()`)

Read locks just exist to check if there's a full lock (someone is writing to it) or not. We can have multiple reads at once, but, we can't read while writing. When `Lock()` is done, all read locks, and regular locks must be cleared, before it can run. However, when `RLock()` is done, only the regular locks need to be cleared.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1","id2","id3","id4","id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1) // increments 1 to the counter
		go dbCall(i)
	}
	wg.Wait() // waits for the counter to go back to 0
	fmt.Println(time.Since(t0))
}

func dbCall(i int) {
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println(dbData[i])
	save(dbData[i])
	log()
	wg.Done() // decrements 1 from the counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println(results)
	m.RUnlock()
}
```

---
# Channels

1. Channels hold data
2. They are thread safe. They avoid data races
3. They listen for data, and we can block code execution when data is added or removed from a channel

```go
package main

import "fmt"

func main() {
	var c = make(chan int) // this channel can only hold a single int value
	c <- 1 // write 1
	var i = <- c // read from the channel
	fmt.Println(i)
}
```

This will cause a **DEADLOCK** !!!

As the channel waits for someone to read when it writes something in. 

```go
package main
import "fmt"

func main() {
	var c = make(chan int)
	go process(c)
	fmt.Println(<-c)
}

func process(c chan int) {
	c <- 123
}
```

**For Loops** works perfectly with channels too

```go
package main
import "fmt"

func main() {
	var c = make(chan int)
	go process(c)
	for i := range c { // it will wait for something to enter the channel.
	// also, it won't close, unless we use the close(c) command in the other func
		fmt.Println(<-c)
	}
}

func process(c chan int) {
	// defer runs the statement just before the function ends 
	defer close(c) // close causes the channel to end, preventing a deadlock
	for i:=0; i<5; i++ {
		c <- i
	}
}
```

**Buffered Channels**

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_BURGER_PRICE float32 = 5

func main() {
	var burgerChannel = make(chan string)
	var websites = []string{"burgerking.com", "mcdonalds.com"}
	for i:=range websites {
		go checkBurgerPrices(website[i], burgerChannel)
	}
	sendMessage(burgerChannel)
}

func checkBurgerPrices(website string, burgerChannel chan string) {
	for {
		time.Sleep(time.Second*1)
		var burgerPrice = rand.Float32()*20
		if burgerPrice<=MAX_BURGER_PRICE {
			burgerChannel <- website
			break
		}
	}
}

func sendMessage(burgerChannel chan string) {
	fmt.Println(<-burgerChannel)
}
```


ight, im really confused. I'll have to understand channels properly later.

Topic left: generics

---
[^1]:[Learn Go](https://www.youtube.com/watch?v=8uiZC0l4Ajw) 