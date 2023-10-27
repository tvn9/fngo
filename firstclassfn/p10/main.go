// Higher Order Function - Pratice 1
package main

import "fmt"

type F func() string

func A() string {
	return "hello "
}

func B(a F, s string) string {
	if s != "" {
		return a() + s
	}
	return a() + "world!"
}

func main() {
	fmt.Println(B(A, "Thanh!"))
}
