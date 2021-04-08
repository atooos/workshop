package main

func main() {

}

type User struct {
	FirstName string
	LastName  string
}

// 1- create the function constructor for the User, name it idiomatically.
// add params to this function so the user struct can be initialized zith all its fields.
func NewUser(fn, ln string) *User {
	return &User{
		FirstName: fn,
		LastName:  ln,
	}
}

// 2- create a method SayName attached to User struct
// this function return a string with the concatenated values "Hello "+ the user frist name value.
func (u *User) SayName() string {
	return "Hello " + u.FirstName
}

// 3- create a method FName that has a param fn in string and mutate the user FristName field.
func (u *User) Fname(fn string) {
	u.FirstName = fn
}
