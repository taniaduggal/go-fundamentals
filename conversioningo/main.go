package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the pizza app")
	fmt.Println("Please rate us")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim any extra spaces or newline characters
	input = strings.TrimSpace(input)

	fmt.Println("Thanks for rating,", input)

	// Debugging: Print the trimmed input
	fmt.Printf("Trimmed input: '%s'\n", input)

	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Add 1 to your rating:", numRating+1)
	}
}
