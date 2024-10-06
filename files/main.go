package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")
	content := "This needs to go in a file - LCO"
	file, err := os.Create("./mylcogofile.txt")
	// if err != nil {
	// 	panic(err) // it shutdown the execution of programand show you what the error you are facing - common way of handling error
	// }
	checkNilErr(err)
	//for writing into in file, we use io package
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

//readingthefile

func readFile(filename string) {
	//writing- there is io utility but for reading and for little-bit more manipulations there is ioutils package
	//databyte, err := ioutil.ReadFile(filename) ioutil pacakge is deprected from go1.16
	databyte, err := os.ReadFile(filename)
	//when ioutil package read the file, its not read in string format, also when you read data from inteneet, that is radd in bytes, so same here
	checkNilErr(err)
	fmt.Println("the bytes are: ", string(databyte))

}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
