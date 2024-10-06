package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	increment := counter()
	fmt.Println(increment()) // prints 1
	fmt.Println(increment()) // prints 2
}

// closures - bahr or ou scope ke values/variable use kre inside the function and it always remember the that variable even after the call stack
