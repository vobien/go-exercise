package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (user user) toString() {
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt.Format("02/01/2006 03:04:05 PM"))
}

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthdate := getUserData("Enter your birthdate (example 12/31/1991): ")

	user := user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

	// solution 1
	user.toString()

	// solution 2
	printOutput(user)
}

func printOutput(user user) {
	// Time formatting with Layout
	// The month comes first at 01, followed by the day of the month at 02,
	// then the hour at 03, the minute at 04, the second at 05, the year at 06 (or 2006) and,
	// finally, the time zone at 07.
	fmt.Printf("Your name is %s %s and your birthdate is %s, your account is created at %s\n",
		user.firstName, user.lastName, user.birthdate, user.createdAt.Format("02/01/2006 15:04:05"))
}

func getUserData(title string) string {
	fmt.Print(title)
	var userData string
	fmt.Scan(&userData)
	return userData
}
