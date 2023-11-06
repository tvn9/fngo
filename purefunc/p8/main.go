// Example function to create user
package main

type User struct {
	username string
	password string
}

type userDb struct {
	username string
	password string
}

/* Example of a function that try to handle too many tasks,
	the function could be broken down into multiple functions
func createUser(username, password string) {
	u := User{
		username,
		password,
	}
	if u.validePassword() {
		userDB.save(u)
	} else {
		panic("invalid password")
	}
}
*/

func signup(username, password string) {
	user, err := createUser(username, password)
	if err != nil {
		saveUser(user)
	} else {
		panic("could not create account")
	}
}

func createUser(username, password string) (User, error) {
	u := User{username, password}
	if u.validPassword() {
		return u, nil
	}
	return User{}, Error.new("invalid password")
}

func (u User) validPassword() bool {
	p := userDb{}
	if u.password == p.password {
		return true
	}
	return false
}

func saveUser(u User) {
	userDb.save(u)
}

func main() {

}
