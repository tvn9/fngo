// Function inside structs
package main

import "fmt"

type predicate func(i int) bool

type ConstraintChecker struct {
	largerThan  predicate
	smallerThan predicate
}

func (c ConstraintChecker) check(input int) bool {
	return c.largerThan(input) && c.smallerThan(input)
}

func main() {
	checker := ConstraintChecker{
		largerThan: createLargerThanPredicate(2),
		smallerThan: func(i int) bool {
			return i < 10
		},
	}
	fmt.Printf("%v\n", checker.check(3))
}

func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}
