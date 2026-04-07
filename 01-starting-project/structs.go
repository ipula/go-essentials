package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin, _ := user.NewAdmin("test@mailinator.com", "password123")
	admin.User.OutputUserData()
	admin.User.ClearUserName()
	admin.User.OutputUserData()

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
