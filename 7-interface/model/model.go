package model

type User struct {
	firstName string
	lastName  string
}

func (u *User) SetFirstName(fn string) {
	u.firstName = fn
}

func (u *User) SetLastName(ln string) {
	u.lastName = ln
}

type Admin struct {
	firstName string
	lastName  string
}

func (a *Admin) SetFirstName(fn string) {
	a.firstName = fn
}

func (a *Admin) SetLastName(ln string) {
	a.lastName = ln
}
