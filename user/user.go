// Groups data and functions into collections
package user

import (
	"fmt"
	"time"
)

type UserType struct {
	firstName   string
	lastName    string
	dateOfBirth string
	age         int64
	createdAt   time.Time
}

type Admin struct {
	email    string
	password string
	User     UserType
}

// Method (a function binded to a struct is termed as method)

// Param is passed as a copy of the userStruct instance not the actual
// value that is why don't miss the * to get it as pass by reference assignment.
func (userInfo *UserType) DisplayUserData() {
	fmt.Println(*userInfo)
}

func (userInfo *UserType) ClearUserData() {
	userInfo.firstName = ""
	userInfo.lastName = ""
	userInfo.dateOfBirth = ""
}

func NewUser(firstName, lastName, dateOfBirth string) (*UserType, error) {
	return &UserType{
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		age:         32,
		createdAt:   time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: UserType{
			firstName: "Admin",
			lastName:  "",
			createdAt: time.Now(),
		},
	}
}
