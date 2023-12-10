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

```
type predicate func(i int) bool

```  
Example of how to use predicate function 

```
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

example of how to pass function to function  

```
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

