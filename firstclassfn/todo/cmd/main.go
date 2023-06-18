package main

import "github.com/tvn9/fpgo/todo/pkg"

func main() {
	t := pkg.NewTodo()
	t.Write("hello world")
}
