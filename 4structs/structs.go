package main

import (
	"bufio"

	"fmt"
	"os"
	"runtime"
	"strings"

	"example.com/example/4structs/user"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (DD/MM/YYYY): ") //changed unchecked

	// ... do something awesome with that gathered data!

	var appUser *user.User

	appUser, err := user.NewUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return //ends app
	}

	admin := user.NewAdmin("test@example.com", "password")

	admin.User.OutputUserDetails()
	admin.User.ClearUsername()
	admin.User.OutputUserDetails()

	// fmt.Println(firstName, lastName, birthdate)
	// outputUserDetails(appUser)
	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
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
