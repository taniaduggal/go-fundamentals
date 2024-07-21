package main

import "fmt"

func main() {
	// const bahr bhi declare kr skte
	// shorthand variable bahr declare nhi hoyega

	const name = "tania"

	fmt.Println(name)
	//const grouping
	const (
		a = 100
		b = "localhost"
		c = 3
	)
	fmt.Println(a, b, c)
}
