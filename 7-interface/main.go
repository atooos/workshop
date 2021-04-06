package main

func main() {
	// implement the interface Changer by adding all methods set to the User struct.
	// the functions SetFirstName and SetLastName should respectivly mutate the user firstName and lastName.
	// initiate a user with the fields
	// user one of the function and display the reslut.
}

type User struct {
	firstName string
	lastName  string
}

type Changer interface {
	SetFirstName(fn string)
	SetLastName(fn string)
}
