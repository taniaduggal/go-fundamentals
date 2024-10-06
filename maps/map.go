package main

import (
	"fmt"
	"maps"
)

func main() {
	// languages := make(map[string]string)

	// languages["JS"] = "Javascript"
	// languages["RB"] = "Ruby"
	// languages["PY"] = "Python"

	// fmt.Println("list of all langs: ", languages)
	// fmt.Println("js shorts for: ", languages["JS"])

	// delete(languages, "RB")
	// //clear(languages) this is inbuilt function clears the map
	// // if key ot exists in the map and we want to access that then it returns the zeroed value of value in map
	// fmt.Println("list of all langs: ", languages)

	// //looping over maps
	// for key, value := range languages {
	// 	fmt.Printf("For key %v, value is %v\n", key, value)
	// }
	// //when u dont need the first value
	// for _, value := range languages {
	// 	fmt.Printf("For key %v, value is %v\n", value)
	// }

	m := map[string]int{"price": 40, "phones": 3}
	fmt.Println(m)

	//we always do work to find elements in the map
	v, ok := m["phones"] // v returns value and ok returns the bool value(true or false) if value exists
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	//comparison of mpsd, we can do with the help of map package like slice, we cant do with ++ sign

	m1 := map[string]int{"price": 40, "phone": 3}
	m2 := map[string]int{"price": 40, "phone": 3}
	fmt.Println(maps.Equal(m1, m2))
}
