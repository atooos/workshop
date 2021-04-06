package main

import "time"

func main() {
	// 1- Compose User with the DBFields and OptionalFields structs.

	// 2- Create a constructor function for User struct that initialize all fields
	// and use the optional fields (OptionalFields) in params withn a pointer value.

	// 4- Use this current function constructor in order to create a new User.

	// 5- Display the result with fmt.
}

type User struct {
	FirstName string
	LastName  string
}

type DBFields struct {
	ID         string
	CreatDate  time.Time
	UpdateDate time.Time
	DeleteDate time.Time
}

type OptionalFields struct {
	Email          string
	EmailValidated bool
	IsConnected    bool
}
