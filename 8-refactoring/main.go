package main

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

func main() {
	// 1 - create a folder and name it model.
	// 2 - move the Pet model and all it function's to a separated file model/pet.go.
	// 3 - do the same for User struct in a file model/user.go.
	// 4 - call the NewUser function in order to create a new user variable called u.
	// 5 - use the function fmt.Println to display the variable value in u.
	// 6 - use the method of u SayHello.
	// 7 - create tow Pets with it's constructor and name it p1 and p2.
	// 8 - add this new pets to u with the AddPet function.
	// 9 - use the function fmt.Println to display u.
}

type Pet struct {
	ID          string
	Name        string
	Description string
}

func NewPet(name, desc string) *Pet {
	return &Pet{
		ID:          uuid.New().String(),
		Name:        name,
		Description: desc,
	}
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int
	Pets      []Pet
}

func NewUser(fn, ln string, age int) *User {
	return &User{
		ID:        uuid.New().String(),
		FirstName: fn,
		LastName:  ln,
		Age:       age,
	}
}

func (u *User) SayHello(msg string) string {
	return u.FirstName + " " + strings.ToUpper(u.LastName) + ": " + msg
}

func (u *User) AddPet(p *Pet) error {
	if p == nil {
		return errors.New("no pets to add")
	}
	u.Pets = append(u.Pets, *p)
	return nil
}

func (u *User) String() string {
	if len(u.Pets) == 0 {
		return "Hello my name is " + u.FirstName
	}
	res := "Hi I'm " + u.FirstName + " and let me introduce you to "
	for _, p := range u.Pets {
		res += p.Name + ", "
	}
	return res[:len(res)-2]
}
