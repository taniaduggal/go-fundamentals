package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums)

	//array length
	fmt.Println(len(nums))
	//zero values
	//to declare in single line
	// num := [3]int{1, 2, 3}

	//2d arrays
	// nums := [2][2]int{{1, 2}, {3, 4}}
	// fmt.Println(nums)
	// why to use
	// -1. fixed size, that is predictable, memoey optimzation, constant time ccess
}
