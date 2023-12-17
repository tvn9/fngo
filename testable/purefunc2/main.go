// Example of functions that create new user
package main

import (
	"errors"
	"fmt"
	"log"
)

type (
	// Name, and Password of type string
	Name     string
	Password string
)

// User holds user name and password
type User struct {
	Name     Name
	Password Password
}

// userDb holds users of type []User
type userDb struct {
	users []User
}

// db declaration
var db userDb

// signup start the new user signup, calling createUser, and saveUser functions
func signup(name Name, password Password) {
	user, err := createUser(name, password)
	if err != nil {
		log.Fatal(err)
	}
	db.saveUser(user)
}

// createUser create new user
func createUser(name Name, password Password) (User, error) {
	if ValidPassword(name, password) {
		u := User{name, password}
		return u, nil
	} else {
		return User{}, errors.New("invalid user name or password")
	}
}

// ValidPassword validate user name and password for correct pattern
func ValidPassword(n Name, p Password) bool {
	name := n
	passwd := p
	if n == name && p == passwd {
		return true
	} else {
		return false
	}
}

// saveUser saves the new user to database
func (u *userDb) saveUser(user User) {
	u.users = append(u.users, user)
}

// PrintDB prints the database content to the screen
func (u userDb) PrintDB() {
	fmt.Println(u)
}

// the beginning
func main() {
	var (
		name1     Name     = "Thanh"
		name2     Name     = "Thanh2"
		name3     Name     = "Thanh3"
		password1 Password = "password"
		password2 Password = "password2"
		password3 Password = "password3"
	)
	signup(name1, password1)
	signup(name2, password2)
	signup(name3, password3)

	db.PrintDB()
}
