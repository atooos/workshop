package main

import (
	"fmt"
	"workshop/7-interface/model"
)

func main() {
	// implement the interface Changer by adding all methods set to the User struct.
	// the functions SetFirstName and SetLastName should respectivly mutate the user firstName and lastName.
	// initiate a user with the fields
	var u model.User
	u.SetFirstName("Donal")
	u.SetLastName("Duck")
	// user one of the function and display the reslut.
	fmt.Println(u)

	var a model.Admin

	var tbl []Changer
	tbl = append(tbl, &u, &a)
	fmt.Println(tbl)
}

type Changer interface {
	SetFirstName(fn string)
	SetLastName(ln string)
}
