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
		largerThan: createLargerThanPredicate(5),
		smallerThan: func(i int) bool {
			return i < 10
		},
	}

	checker1 := ConstraintChecker{
		largerThan:  createLargerThanPredicate(2),
		smallerThan: createSmallerThanPredicate(5),
	}
	fmt.Printf("%v\n", checker.check(6))
	fmt.Printf("%v\n", checker1.check(3))
}

func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}

func createSmallerThanPredicate(threadhold int) predicate {
	return func(i int) bool {
		return i < threadhold
	}
}
