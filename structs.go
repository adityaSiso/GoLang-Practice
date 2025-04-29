// Groups data and functions into collections
package main

import (
	"fmt"

	"example.com/go-module/user"
)

func getUserIntput(inputText string) string {
	var userInput string
	fmt.Print(inputText)
	fmt.Scan(&userInput)
	return userInput
}

func getUserData() {
	firstName := getUserIntput("First Name: ")
	lastName := getUserIntput("Last Name: ")
	dateOfBirth := getUserIntput("Date of Birth in DD/MM/YYYY format: ")

	var userStruct *user.UserType

	userStruct, _ = user.NewUser(firstName, lastName, dateOfBirth)
	userStruct.DisplayUserData()
	userStruct.ClearUserData()
	userStruct.DisplayUserData()
}
