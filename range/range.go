package main

import "fmt"

//iterating over data structures

func main() {
	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums), i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
	// for i, value := range nums ---> i returns index(if we iterate through slices) and value returns value and if we dont want or ignore it, simply written like this _

	//iteration over map
	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
		// in map,range with for, for ives first as key and second is value.
	}

	//range can be used upon string
	// i gives index and c gives unicode  (unicode code point rune)
	// i is basically starting byte of rune
	// very imp concept here , when u revis the code, you can again have alook
	for i, c := range "hello" {
		fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
