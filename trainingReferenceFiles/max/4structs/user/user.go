package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User //  called embedding, but basically extending User
}

// user here is called a receiver, indicating the struct to which it'll be attached
func (u User) OutputUserDetails() { //attaches as method to user struct
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUsername() { //must be a pointer or it'll mutate a copy
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN", //could have been passed, but dummy vals for simplicity
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// Constructor function is a pattern rather than lang feature
// in external packages, it's commonly just New as it'll import as User.New()
func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("No blank fields allowed")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil

}
