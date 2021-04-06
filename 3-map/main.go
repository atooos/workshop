package main

import "fmt"

func main() {

	// 1- fix the delared map and change from the key b the value into B
	var ma map[string]string = map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
		"d": "D",
	}
	// 2- display the current map values.
	fmt.Println("ma", ma)

	// 3- create a map called mu of map of string of a User
	// create a key called e15e98dc-33ca-11eb-b206-1f97348a99b3 and insert a User and display it.
	mu := make(map[string]*User)
	mu["e15e98dc-33ca-11eb-b206-1f97348a99b3"] = &User{
		FirstName: "Rob",
		LastName:  "Pike",
	}
	fmt.Println("user", mu["e15e98dc-33ca-11eb-b206-1f97348a99b3"])

	// 4- (advenced) delete the current created key
	// and display if there is something into this current key.
	delete(mu, "e15e98dc-33ca-11eb-b206-1f97348a99b3")
	u, ok := mu["e15e98dc-33ca-11eb-b206-1f97348a99b3"]
	if u == nil {
		fmt.Println("no user inside the map", u)
	}
	fmt.Println("is there something ?", ok)
}

type User struct {
	FirstName string
	LastName  string
}
