package main

import "fmt"

func main() {
	// if-else
	// age := 18
	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }

	// else-if
	// age := 18
	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 15 && age < 18 {
	// 	fmt.Println("person is a teenager")
	// } else {
	// 	fmt.Println("person is a child")
	// }

	// var role = "admin"
	// var hasPermissions = false

	// if role == "admin" && hasPermissions {
	// 	fmt.Println("yes")
	// }

	//one more syntax of if that we can declare variable in if statement
	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is a teenager", age)
	} else {
		fmt.Println("person is a child", age)
	}

	// go doesnt have ternary operator, u will have to use normal if-else
}
