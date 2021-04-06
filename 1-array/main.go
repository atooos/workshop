package main

import "fmt"

func main() {
	// 1- fix this bug so the array has 4 elements.
	var tbl1 [4]string = [4]string{
		"1", "2", "3", "4",
	}

	// 2- fix the bug so the array contain string.
	var tbl2 [8]string = [8]string{
		"z", "b", "c", "d",
	}

	// 3- tbl1 : change the value at index 0 to a and display the result.
	tbl1[0] = "a"
	fmt.Println(tbl1)

	// 4- copy all the initialized values from tbl2 into tbl1 and display the values of tbl1.
	for i := 0; i < len(tbl1); i++ {
		tbl1[i] = tbl2[i]
	}
	fmt.Println(tbl1)

}
