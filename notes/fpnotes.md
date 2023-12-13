## Nodes and Example Code  

### Type alias for function

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

### Passing function to function  

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

### Inline function  

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

### Anonymous Function  

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

``` go
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























