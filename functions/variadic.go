package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
} // func sum(nums ...interface{}) interface{} = go mein any type ka hota hai
func main() {

	nums := []int{3, 4, 5, 6}
	// result := sum(3, 4, 5)
	result := sum(nums...)
	fmt.Println(result)
}
