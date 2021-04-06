package main

import "fmt"

func main() {
	// 1- fix the bug so the slice is valid slice in the declaration.
	var tbl1 []string = []string{
		"z", "b", "c", "d",
	}

	// 3- tbl1: display the capacity, lenght and all it's values.
	fmt.Printf("len :%v / cap :%v\n", len(tbl1), cap(tbl1))

	// 4- copy all values from the tbl1 into a new slice called tbl2
	// and display values of new tbl2 slice.
	tbl2 := make([]string, len(tbl1))
	copy(tbl2, tbl1)
	fmt.Println("tbl2:", tbl2)

	// 5- change a value from tbl1 and display all values from tbl2,
	// if no values has changed into tbl2,
	// you succeed to create a copy and not a slice of it :)
	tbl1[0] = "a"
	fmt.Println(tbl2[0] == "a")

	// 6- display the capacity and the lenght of tbl2
	// use append to add tow elements into tbl2
	// display again the capacity and lenght of tbl2.
	fmt.Printf("len :%v / cap :%v\n", len(tbl2), cap(tbl2))
	tbl2 = append(tbl2, "e", "f")
	fmt.Printf("len :%v / cap :%v\n", len(tbl2), cap(tbl2))

	// 7- (advanced) add 2000 random string into tbl2
	// at each new item added, check if the capacity has changed
	// display the new capacity.
	oldCap := cap(tbl2)
	for i := 0; i < 2000; i++ {
		tbl2 = append(tbl2, string(rune(i)))
		if oldCap != cap(tbl2) {
			oldCap = cap(tbl2)
			fmt.Println(oldCap)
		}
	}
}
