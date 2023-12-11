# FUNCTIONAL PROGRAMMING IN GO

Practice functional programming in Go  
The goal of this coding practice is to gain mastery on how to use functional programming concept in Go.

## Resources and Tools  

### Book:  
The Go Programming Language - Alan A.A. Donovan - Brian W. Kerninghan  
Functional Programming in Go - Dylan Meeus

### Tools:  
Go Standard Library  
VS Code 

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
   // Filter function uses an un named slice and function to look for even numbers
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



