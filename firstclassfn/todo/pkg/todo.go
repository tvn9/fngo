package pkg

import "log"

// Todo truct
type Todo struct {
	Text string
	Db   *Db
}

// NewTodo create and return a Todo object
func NewTodo() Todo {
	return Todo{
		Text: "",
		Db:   NewDB(),
	}
}

// Write write string to ToDo's Text
func (t *Todo) Write(s string) {
	if t.Db.IsAuthorized() {
		t.Text = s
	} else {
		log.Fatal("user not authorized to write")
	}
}

// Append adds text to the existing Todo's text field
func (t *Todo) Append(s string) {
	if t.Db.IsAuthorized() {
		t.Text += s
	} else {
		log.Fatal("user not authorized to append")
	}
}
