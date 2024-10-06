// // alternative of classes..

// methods, instances, constructors, inline-struct

// package main

// import (
// 	"fmt"
// 	"time"
// )

// //there is no iheritance and super parents like things here

// // func main() {
// // 	fmt.Println("structs in golang")
// // 	tania := User{"Tania", "tsnisduggsl@gamil.com", true, 19}
// // 	fmt.Println(tania)
// // 	fmt.Printf("tania details are: %+v\n", tania) //+v provides more details
// // 	fmt.Printf("Name is %v and email is %v\n", tania.name, tania.Email)
// // }

// // // it should be in uppercase that we can import it easily
// // type User struct {
// // 	name   string
// // 	Email  string
// // 	Status bool
// // 	Age    int
// // }

// //struct is a way to recognixse oops in go
// //clases ko kuch methods attach krte hain

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanaosecond precision
// }

// // reeciver - by convention we use first letter of struct
// func (o *order) changeStatus(status string) {
// 	o.status = status // idr struct automaticalllly hmare liye dereferncing ka kam krta hai, star dene ki need nahi hai
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

// func main() {

// 	//inline struct

// 	language := struct {
// 		name   string
// 		isGood bool
// 	}{"golang", true}

// 	myOrder := newOrder("1", 30.50, "RECIVED")
// 	fmt.Println(myOrder.amount)
// 	// myOrder := order{
// 	// 	id:     "1",
// 	// 	amount: 50.00,
// 	// 	status: "received",
// 	// }
// 	//if you dont set any field toh its gives zero values
// 	// myOrder.changeStatus("confirmed")
// 	// fmt.Println(myOrder)
// 	// fmt.Println(myOrder.getAmount())

// 	// jb hum struct define krte toh sath hi kuch initial chije setup krne ki need hoti hai, that thing we can do with constructors. jb construt bnega the values we defined in the main function for passing, voh hum constructor mein likhenge. in c, constructor class ki properties ko initialize keta hai

// }

// //ideally therese no constructor in go,  but we are doing similar to amke constructor. covention to make constructor in go , write with new keyword

// func newOrder(id string, amount float32, status string) *order {
// 	//INITIAL SETUP GOES HERE
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// 	//one more convention, we never return struct directly, we have to return pointer of it
// }
