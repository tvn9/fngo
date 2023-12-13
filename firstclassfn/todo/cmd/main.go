package main

import (
	"fmt"

	"github.com/tvn9/fpgo/firstclassfn/todo/pkg"
)

func main() {
	t := pkg.NewTodo()

	t.Write("Hello, world!")

	fmt.Println(t.Text)
}
