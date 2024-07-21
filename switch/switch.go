package main

import (
	"fmt"
)

func main() {
	// i := 3
	// //simple switch
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// 	//no need to write break, it automatically handles behind the scene
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default: //optional
	// 	fmt.Println("other")
	// }

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its a weekend")
	// default:
	// 	fmt.Println("its a weekday")
	// }
	// type switch
	whoAmI := func(i interface{}) {
		// when there is nothing in interafce brackets, it can recive any value
		switch t := i.(type) {
		// switch i.(type)
		// type identifies the type of i and save it itno t

		case int:
			fmt.Println("its a integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("its a different type", t)
		}
	}
	whoAmI(1.1)

}
