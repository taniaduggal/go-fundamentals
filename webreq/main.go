package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://gobyexample.com/"

func main() {
	fmt.Println("demo web req")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response is: ", response)
	defer response.Body.Close() // caller's responsibilty to close the connection
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
