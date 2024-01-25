package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// user here is called a receiver, indicating the struct to which it'll be attached
func (u user) outputUserDetails() { //attaches as method to user struct
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUsername() { //must be a pointer or it'll mutate a copy
	u.firstName = ""
	u.lastName = ""
}

//Constructor function is a pattern rather than lang feature

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ") //changed unchecked

	// ... do something awesome with that gathered data!

	var appUser *user

	appUser = newUser(
		userFirstName,
		userLastName,
		userBirthdate)

	// fmt.Println(firstName, lastName, birthdate)
	// outputUserDetails(appUser)
	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
}

// func outputUserDetails(u user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	var cleanedInput string

	if runtime.GOOS == "windows" {
		cleanedInput = strings.Replace(userInput, "\r\n", "", -1)
	} else {
		cleanedInput = strings.Replace(userInput, "\n", "", -1)
	}
	return cleanedInput

}
