//introduces in go 1.18+
// to solve do not try yourself(duplicacy)- we use generics

package main

import "fmt"

// T is a convention, andis of any type. it becomes a generic
// func printSlice[T any/interface{}](items []T), interace{} = any. its not a good practice because if we set any, toh koi bhi type isme aa skti. toh we can scope it for some types only

// func printSlice[T int | string | bool](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	} // generic on function
// }

// we can use comparable type even because we have more types to comapare reather than only above.
// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	} // generic on function
// }

// we can pass more types
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	} // generic on function
}

// func printSlice[T comparable](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	} // generic on function
// }

//generic on struct
//LIFO
// type stack struct {
// 	elements []int
// }

type stack[T any] struct {
	elements []T
}

// func printSlice[T any](items []T) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	} //duplicacy of code
// }

func main() {
	// myStack := stack[string]{
	// 	elements: []string{"golang"},
	// }
	// fmt.Println(myStack)
	// nums := []int{1, 2, 3}
	// printSlice(nums)
	// printSlice([]int{1, 2, 3})
	// 	names := []string{"golang", "ts"}
	// 	printStringSlice(names)
	vals := []bool{true, false, true}
	printSlice(vals, "john")
}
