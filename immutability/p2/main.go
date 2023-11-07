// go immutable in slice
package main

import "fmt"

func main() {
	names := []string{"Bernie"}
	names = append(names, "Kim")
	names = append(names, "Thanh")

	fmt.Printf("%v\n", names)

	// When passing a slice to func to add value, the newly added
	// value will only be on the copy slice
	newNames := addValue(names, "Wendy")
	fmt.Printf("The copy newNames slice = %v\n", newNames)

	fmt.Printf("The original names slice = %v\n", names)

}

func addValue(s []string, name string) []string {
	s = append(s, name)
	return s
}
