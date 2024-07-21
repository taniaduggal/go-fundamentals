// only for construct is available in go for ooping

package main

import "fmt"

func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	//infinite loop
	// for {
	// 	println("1")
	// }

	//classic for loop
	for i := 0; i <= 3; i++ {
		// break - stop the loop
		// continue- skip the current iteration
		// fmt.Println(i)

		if i == 2 {
			continue
		}
		fmt.Println(i)

		//1.22 range- kuch chij kuch time ke liye krni hai toh we can use range
		for i := range 3 {
			// i starts from 0 and goes upto 2 or before 3
			fmt.Println(i)
		}

	}
}
