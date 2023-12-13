package pkg

import (
	"fmt"
	"os"
)

// type alias for function
type authorizationFunc func() bool

// Db struct
type Db struct {
	AuthorizationFn authorizationFunc
}

// IsAuthorized method
func (d *Db) IsAuthorized() bool {
	return d.AuthorizationFn()
}

func NewDB() *Db {
	return &Db{
		AuthorizationFn: argsAuthorization,
	}
}

func argsAuthorization() bool {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("usage: comamand username")
		os.Exit(1)
	}
	user := "admin"
	// super secure authorization layer
	// in a real application, this would be a database call
	fmt.Println(user)
	return args[0] == user
}
