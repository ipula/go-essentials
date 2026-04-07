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
	User     User
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func (u User) OutputUserData() {
	// if using pointers (*u).firstName
	fmt.Printf("User: %s %s\n", u.firstName, u.lastName)
	fmt.Printf("Birthdate: %s\n", u.birthdate)
	fmt.Printf("Account created at: %s\n", u.createdAt.Format(time.RFC1123))
}

func NewAdmin(email, password string) (Admin, error) {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "User",
			birthdate: "N/A",
			createdAt: time.Now(),
		},
	}, nil
}

// constructor function for user struct
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		fmt.Println("Error: All fields are required to create a user.")
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
