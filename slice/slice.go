package main

import (
	"fmt"
)

func main() {
	// uninitializes slice is nil
	// var nums []int
	// var nums = make([]int, 2, 5)
	// nums[0] = 3
	// make also initilze
	// make(initilization, length, capacity)
	// capacity = maximum number of elements can fit
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)

	//other way to make a  slice
	// nums := []int{} //creates an empty slice
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	//cap changes or resize by multiply by 2

	//copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums) // copy(dest, src)
	// fmt.Println(nums, nums2)

	// slice operator is :
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2]) // it excludes the last index
	// fmt.Println(nums[:2])  // from to value
	// fmt.Println(nums[1:])  // now last value is till end, not exclude anything

	// slice inbuilt package
	// comapring
	// var nums1 = []int{1, 2, 3}
	// var nums2 = []int{1, 2, 3}
	// fmt.Println(slices.Equal(nums1, nums2)) // it returns bool value

	// 2d slices
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}
