## Nodes and Example Code  

### Type alias for functions

``` go
// type alias for function
type predicate func(i int) bool

```  
Example of how to use predicate function 

``` go
// Filter function accept a number slice, and a func alias
// that take an int and return bool.
func Filter(nums []int, p predicate) []int {
   var out []int
   for _, n := nums {
      // calling the predicate function
      if p(n) {
         out = append(out, n)
      }
   }
   return out
}
``` 

### Passing functions to functions  

Example of how to pass function to function  

```go
// type alias for function
type predicate func(i int) bool  

// evenNum function returns true or false when i is an even number
func evenNum(i int) bool {
   return i % 2 == 0 
}

// Filter function returns number slice that contains all even number
func Filter(nums []int, p predicate) []int {
   var out []int
   for _, n := range nums {
      if evenNum(n) {
         out := append(out, n)
      }
   }
   return out
}
```  

### Inline functions  

```go
func main() {
   nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
   // evenNum return true or false when n is even or odd
   evenNum := func(n int) bool {
      return n % 2 == 0 
   }
   fmt.Println(Filter(nums, evenNum))
}
```

### Anonymous functions 

```go
// anonymous function
Filter(nums, func(i int) bool { return i % 2 == 0 })
```

Example of how to use and anonymous function 

```go
func main() {
   // Filter function uses an unnamed slice and function to look for even numbers
   out := Filter([]int{2,3,5,8,10}, func(i int) bool { return i % 2 == 0 })
   fmt.Println(out)
}
```  

### Return function from function 

``` go
// type alias for function
type predicate func(i int) bool  

// largerThan allows the calling function to pass the threadhold value
func largerThan(filter int) predicate {
   return func(i) bool { return i > filter}
}

// smallerThan allows the calling function to pass the threadhold value
func smallerThan(filter int) predicate {
   return func(i int) bool { return i < filter }
}
```  

Example of how to use a return function from function

``` go 
// function in var 
var ( 
   largerThanFive = largerThan(5)
   largetThan50 = largerThan(50)
   smallerThanTwo = smallerThan(2)
)

func main() {
   nums := []int{1, 2, 5, 8, 10, 23, 59, 99, 100}

   larger5 := Filter(nums, largerThanFive) // [8, 10, 23, 99, 100]
   larget50 := Filter(nums, largerThan50)  // [59, 99, 100]
   smaller2 := Filter(nums, smallerThanTwo) // [1]

   fmt.Println(larger5)
   fmt.Println(larger50)
   fmt.Println(smaller2)
} 
```

### Function inside data structures

Example of function in a slice of functions - []func  

```Go
predicates := []predicates{function1, function2, function3}
```  

How to use function in slice of functions

``` go 
// function in var
var (
   largerThanFive = largerThan(5)
   largetThan50 = largerThan(50)
   smallerThanTwo = smallerThan(2)
)

func main() {
   nums := []int{1, 2, 5, 8, 10, 23, 59, 99, 100}
   
   predicates := []predicate{largerThanFive, largerThan50, smallerThanTwo}

   for _, predicate := range predicates {
      out := Filter(nums, predicate)
      fmt.Println(out)
   } 
}
```  

Example of function inside map 

``` go
predicates := map[string]predicate{
   ["larger5": largerThan(5)
   ["larger10": largerThan(10)
   ["smaller23": smallerThan(23)
}
```

How to use functions inside a map  

``` go 
func main() {
   nums := []int{1, 2, 5, 8, 10, 23, 59, 99, 100}
   predicates := [string]predicate{
      "larger5": largerThan(5),
      "larger10": largerThan(10,
      "smaller23": smallerThan(23))
   }

   fmt.Printf("%v\n", Filter(nums, predicates["larger5"]))
   fmt.Printf("%v\n", Filter(nums, predicates["larger10"]))
   fmt.Printf("%v\n", Filter(nums, predicates["larger23"]))
}
```

### Function inside structs  

``` go
// type alias for function
type predicate func(i int) bool 

// largerThanFunc returns type predicate function 
func largerThanFunc(filter int) predicate {
   return func(i int) bool {
      return i > filter
   }
}

// ValidCheck holds predicate function
type ValidRange struct {
   largerThan predicate
   smallerThan predicate
}
```  

How to use function in struct  

```go 
// check method returns ValidRange's fields that hold type predicate function
func (v ValidRange) check(i int) bool {
   return v.largerThan(i) && v.smallerThan(i)
}

func main() {
   // create new ValidRage struct that hold the functions 
   // that check for value larger than 2 and smaller than 9
   checker := ValidRange {
      largerThan: largerThanFunc(2),
      smallerThan: func(i int) bool { return i < 9 }
   }
   // Check to see if 6 is greater than 2 and smaller than 9
   fmt.Printf("%v\n", checker.check(6))   // true
}
```

### Higher-order functions

Basic example of higher-order function  

```go
// Introduction to higher-order functions
func Hello() string {
	return "Hello,"
}

func World() string {
	return Hello() + " world!"
}

func main() {
	fmt.Println(World())
}
```

Higher-order function as inter function  

```Go
// Although the example shows there is no passing function as input, 
// but a function can return a function as output.
// This is a valid characteristic of higher-order functions.

func outerfunc(s string) func() {
	fmt.Println("Outer func: ", s)
	return func() {
		fmt.Println("Inter function")
	}
}

func main() {
	testfunc := outerfunc("Higher-order function")
	testfunc()
}
```

Closure as inner function  

```Go
// Example of closure, a closure is an inner function that use
// a variable introduced in the outer function to perform its work.

func ask() func(string) string {
	s := "What is your name? "
	return func(name string) string {
		return s + name
	}
}

func main() {
	interview := ask()
	name := interview("Thanh Vu Nguyen.")
	fmt.Println(name)
}
```

### Partial application 

Example of Partial application  

```Go
// createGreeting returns a function that also take in a string
// combind the variable and return a string
func createGreeting(g string) func(s string) string {
	return func(name string) string {
		return g + name
	}
}

func main() {
	firstGreeting := createGreeting("Well, hello there ")
	secondGreeting := createGreeting("Hola ")
	fmt.Println(firstGreeting("Thanh"))
	fmt.Println(firstGreeting("Xuan"))
	fmt.Println(secondGreeting("Kim"))
}
```  

[Full code example on partial application](../hiorderfunc/dogspawner/main.go)  

### Function curring, how to reduce n-ary function to unary function  

Example of function curring  

```go 
func F(a, b, c): int {}
```  

The first function, (Fa), takes a argument as input and return F(b) as output, function F(b) takes b as input and return function F(c) as output, function F(c) takes c as input and return and in as final output.  

```Go 
func Fa(a): Fb(b)
func Fb(b): Fc(c)
func Fc(c): int 
```
This is done by leveraging the concept of first-class citizens and higher-order functions.  

Example of function without currying 

```Go
func threeSum(a, b, c int) int {
   return a + b + c
}
```  

With currying  

```go
func threeSumCurried(a int) func(int) func(ing) int {
   return func(b int) func(int) int {
      return func(c int) int {
         return a + b + c
      }
   }
}
```

How to use the currying function?  

```Go
func main() {
   fmt.Println(threeSum(30, 10, 20))   // 60
   fmt.Println(threeSumCurried(20)(30)(10))    // 60
}
``` 

### Pure functions 

What is purity? 
- Does not generate any side effects
- Return the same output when providing the same input (idempotence) 
  
Demonstrating pure versus impure function calls  

```Go 
// add take a, b and return the sum of a and b 
// this is a form a pure function, as it will always return 
// the consitent out when call with the same input
func add(a, b int) int {
   return a + b
}
```  

impure function  

```Go
// rollDice return a random number from 0 to 6
// the output is always unpredictable
func rollDice() int {
   return rand.Intn(6)
}
``` 

### Referential transparentcy  

```Go
// Example we have this calculation
X = 1 + (2 * 2)

// the result is 5. We could have gotten the same result had we replaced
// the multiplication with its result, like below:
X = (1 + 4)
```  

This property is what we mean by referential transparency. All mathematical
operations have this property, and many have leveraged this property when
working out equations in algebra, calculus, or other mathematics classes.  

In a programming language, referential transparency means that a function call
can be replaced with its result.

See this example 
```Go
func main() {
   // call add(10, 5) result = 15 + 20, the total = 35
   fmt.Printf("%v\n", add(20, add(10, 5)))  // 35
   // pass the result of function add, 20 + 15, the total = 35  
   fmt.Printf("%v\n", add(10, 15))          // 35
}
func add(a, b int) int {
   return a + b
}
```  

One more example of the function that break the referential transparency. 

```Go
func main() {
   // time.Now() return the system time at execution time therefore the 
   // result is always changing and unpredictable.
   fmt.Println(time.Now())
}
``` 

### Idempotence 

Idempotence is one of the feature of pure functions, this means that the function will always returns the same result no matter how many times it executed.  

## Pure Functionals Implementation Examaples 

Example of a basin onlinestore written in pure functional style

[Online Store](../fpgo/onliestore/main.go)

### Predicate-based functions 

A predicate is a statement that can be evaluated as either true or false.
The common use case is to filter a set of data into a subset that match a specific condition. For example filter all numbers that are smaller or larger than 10.  

Predicate function type using Go generic

```Go
type Predicate[A any] func(A) bool  
```  

Filter function implementation 

```Go
func Filter(numbers []int) []int {
   nums := []int{}

   for _, n := range numbers {
      if n > 10 {
         nums = append(out, n)
      }
   }
   return nums
}
``` 

This function is can only filter a fix value, which is not very usefull in real program. 

The next example will improve the function to a more dynamic filter value. 

```Go
func Filter(numbers []int, f int) []int {
   nums := n range numbers {
      if n > f {
         nums = append(nus, n)
      }
   }
   return nums
}
```

Filter function in generic form 

```Go
func Filter[A any](input []A, p Predicate[A]) []A {
   output := []A{}

   for _, a := range input {
      if p(a) {
         output = append(output, a)
      }
   } 
   return output
}
``` 


















