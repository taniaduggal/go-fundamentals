package main

import "fmt"

func main() {
	// var ptr *int
	myNumber := 23

	var ptr *int = &myNumber // &= it mean am refrencing
	// var ptr *int = &myNumber // &= it mean am refrencing
	fmt.Println(ptr)
	fmt.Println(*ptr) // defrencing, to access the catual value
	*ptr = *ptr + 2
	fmt.Println(myNumber)

}

// other benefit of pointer: resources save krna
