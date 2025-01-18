package creating_structs

import (
	"fmt"
	"time"
)

type Customer struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u Customer) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func NewCustomer(firstName, lastName, birthdate string) Customer {
	return Customer{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
}