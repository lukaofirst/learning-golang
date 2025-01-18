package embedded_structs

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

type Admin struct {
	Email    string
	Password string
	User
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s %s!", u.FirstName, u.LastName)
}