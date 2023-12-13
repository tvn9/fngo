package main

import (
	"fmt"
	"fpgo/firstclassfn/todo/pkg"
)

func main() {
	t := pkg.NewTodo()

	t.Write("Hello, world!")

	fmt.Println(t.Text)
}
