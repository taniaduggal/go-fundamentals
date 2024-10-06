package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("tan.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err) //panic mtlb ab kuch nahi ho skta program ka
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file name:", fileInfo.IsDir())
	// fmt.Println("file name:", fileInfo.Size())
	// fmt.Println("file name:", fileInfo.Mode())
	// fmt.Println("file name:", fileInfo.ModTime())

	//read file
	// f, err := os.Open("tan.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close() //jab main function exit ho jayegi, ye yaha pr call ho jayega
	// //buffer= temporary space memory mein
	// //make- array of bytes hota hai,
	// buf := make([]byte, 12)
	// length, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	println("data", length, string(buf[i]))
	// }

	// is this a good method to read a file? ofc, no. ye sara dta ko ek bar hi memory mein load krke file read krleta hai which is not that much effecient. if file is small toh its ok but if the file is big then it creates the problem like it occupies all resources etc.
	// data, err := os.ReadFile("tan.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//nodejs mein file ko stream krna bdi files sekhege vo
	//we can read folders

	//dir, err := os.Open(".")
	// dir, err := os.Open("../") //iska mtlb ek folder back ja rha hu
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()
	// //n<=0 then readDir puri files read/return krdega
	// //if its 1 then 1 krega, if 2 then 2 krega
	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	// //create a file
	// f, err := os.Create("tan1.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go")
	// f.WriteString("hi go") //ye append krega
	// //homework- replace

	// bytes := []byte("hilo golang")

	// f.Write(bytes)

	// //read from one file and insert into other file- transfer  with stream fashion (straming fashion)

	// sourceFile, err := os.Open("tan.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("tan1.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()
	// //using streaming with buf.io inbuilt package

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		//extra check , sometimes eof error deta, mtlb file toh end nahi hui, iska mtlb kuch toh error hai file read krne mein
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break // important infinite loop se bahr jane ke liye. agr eof error nahi toh or koi ho skta hai

	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// //last mein data ko flush krna hai, agr koi data bcha hain to

	// writer.Flush()

	// fmt.Println("wriiten to new file successfully")
	//we can even directly copy conetent from one file to other file with the help of copy function(check docs)

	//delete a file

	err := os.Remove("tan1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("FILE DELETED SUCCESFULLY")
}
