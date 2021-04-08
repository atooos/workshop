package main

import (
	"fmt"
	"time"
)

func main() {
	// 1- Compose User with the DBFields and OptionalFields structs.

	// 2- Create a constructor function for User struct that initialize all fields
	// and use the optional fields (OptionalFields) in params withn a pointer value.

	// 4- Use this current function constructor in order to create a new User.
	dbFi := DBFields{
		ID:        "e190fe86-97b5-11eb-8a02-e35d2fd17ab3",
		CreatDate: time.Now(),
	}

	u := NewUser("Rob", "Pike", &dbFi, nil)

	// 5- Display the result with fmt.
	fmt.Println(u)
}

type User struct {
	FirstName string
	LastName  string
	DBFields
	OptionalFields
}

func NewUser(fn, ln string, dbFields *DBFields, optional *OptionalFields) *User {
	var u User
	u.FirstName = fn
	u.LastName = ln

	if dbFields != nil {
		u.ID = dbFields.ID
		u.CreatDate = dbFields.CreatDate
		u.UpdateDate = dbFields.UpdateDate
		u.DeleteDate = dbFields.DeleteDate
	}

	if optional != nil {
		u.Email = optional.Email
		u.EmailValidated = optional.EmailValidated
		u.IsConnected = optional.IsConnected
	}
	return &u
}

type DBFields struct {
	ID         string
	CreatDate  time.Time
	UpdateDate time.Time
	DeleteDate *time.Time
}

type OptionalFields struct {
	Email          string
	EmailValidated bool
	IsConnected    bool
}
