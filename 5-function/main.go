package main

import (
	"fmt"
	"time"
)

func main() {}

// 1- create a Hello function signature that has one param str in string and return a string.
func Hello(str string) string {
	return str
}

// 3 - create a function that has three params of int an return the addition of it.
// optimize the function signature as much as you can.
func Todo(a, b, c int) int {
	return a + b + c
}

// 2- create a Log function signature that has a variadic param in interface{} type.
// Inside the function get the time unix format and concatenate it with the given params.
// Display the result into the standart output.
func Log(elems ...interface{}) {
	res := ""
	for k := range elems {
		res += fmt.Sprint(elems[k])
	}
	fmt.Println(time.Now().Unix(), "-", res)
}

// 4- create a Valid function signature that return another function that return a bool (closure).
func Valid() func() bool {
	return func() bool {
		return true
	}
}

// 5- (advenced) modify this closure function in order to create an email validation.
