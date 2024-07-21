package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all langs: ", languages)
	fmt.Println("js shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list of all langs: ", languages)

	//looping over maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	//when u dont need the first value
	for _, value := range languages {
		fmt.Printf("For key %v, value is %v\n", value)
	}
}
