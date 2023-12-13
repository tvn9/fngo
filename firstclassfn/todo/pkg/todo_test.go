package pkg_test

import (
	"fmt"
	"testing"

	"github.com/tvn9/fpgo/firstclassfn/todo/pkg"
)

func TestTodoWrite(t *testing.T) {
	todo := pkg.Todo{
		Db: &pkg.Db{
			AuthorizationFn: func() bool { return true },
		},
	}

	text1 := "Hello, World!"
	fmt.Println("text1:", text1)
	text2 := " Goodbye, crue world!"
	fmt.Println("text2:", text2)
	text3 := text1 + text2
	fmt.Printf("%s + %s = %s\n", text1, text2, text3)

	todo.Write(text1)
	if todo.Text != text1 {
		t.Errorf("Expected %s but got %v\n", text1, todo.Text)
	}

	todo.Append(text2)
	if todo.Text != text3 {
		t.Errorf("Expected %s but go %v\n", text3, todo.Text)
	}
}
