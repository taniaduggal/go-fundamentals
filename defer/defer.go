package main

import "fmt"

func main() {
	defer fmt.Println("world")
	defer fmt.Println("oNE")
	defer fmt.Println("Two")
	fmt.Println("hyy")
	myDefer()
}

//defer statement executes at the end of the code. defered functions invoked immediately before the surrounding function returns returns{LIFO}

// hyy, 4, 3, 2, 0, 1, two, one, world
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
