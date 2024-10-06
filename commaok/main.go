package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ebter the rating for our pizza:")

	//comma ok syntax

	input, _ := reader.ReadString('\n') //here \n means i want to read until or upto \n comes or new line comes

	fmt.Println("thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input) //whether i input integer, it gives or shows rating as string
}
