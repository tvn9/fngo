# FUNCTIONAL PROGRAMMING IN GO

Practice functional programming in Go

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
type predicate func(i int) bool

```  
Example of how to use predicate function 

``` go
func Filter(nums []int, p predicate) []int {
   var out []int
   for _, n := nums {
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
type predicate func(i int) bool  

func evenNum(i int) bool {
   if i > 0 {
      if i % 2 == 0 {
         return true
      }
   }
}

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
   nums := []int{1,2,3,4,5,6,7,8,9,0}
   // find even num
   evenNum := func(n int) bool {
      if n % 2 == 0 {
         return true
      } 
   }

   fmt.Println(Filter(nums, evenNum))
}
```

### Anonymous Function  

```go
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

```go
type predicate func(i int) bool  


func largerThan(filter int) predicate {
   return func(i) bool { return i > filter}
}

func smallerThan(filter int) predicate {
   return func(i int) bool { return i < filter }
}
```  

Example of how to use a return function from function



