package main

import (
	"fmt"
	"workshop/8-refactoring/model"
)

func main() {
	// 1 - create a folder and name it model.
	// 2 - move the Pet model and all it function's to a separated file model/pet.go.
	// 3 - do the same for User struct in a file model/user.go.
	// 4 - call the NewUser function in order to create a new user variable called u.
	u := model.NewUser("Rob", "Pike", 50)
	// 5 - use the function fmt.Println to display the variable value in u.
	fmt.Println(u)
	// 6 - use the method of u SayHello.
	u.SayHello("world")
	// 7 - create tow Pets with it's constructor and name it p1 and p2.
	p1 := model.NewPet("Mimichat", "this is a cat")
	p2 := model.NewPet("Hector", "this is a dog")
	// 8 - add this new pets to u with the AddPet function.
	u.AddPet(p1)
	u.AddPet(p2)
	// 9 - use the function fmt.Println to display u.
	fmt.Println(u)
}
