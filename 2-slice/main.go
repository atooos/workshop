package main

func main() {
	// 1- fix the bug so the slice is valid slice in the declaration.
	var tbl1 []string = [4]string{
		"z", "b", "c", "d",
	}

	// 3- tbl1: display the capacity, lenght and all it's values.

	// 4- copy all values from the tbl1 into a new slice called tbl2
	// and display values of new tbl2 slice.

	// 5- change a value from tbl1 and display all values from tbl2,
	// if no values has changed into tbl2,
	// you succeed to create a copy and not a slice of it :)

	// 6- display the capacity and the lenght of tbl2
	// use append to add tow elements into tbl2
	// display again the capacity and lenght of tbl2.

	// 7- (advanced) add 2000 random string into tbl2
	// at each new item added, check if the capacity has changed
	// display the new capacity.

}
