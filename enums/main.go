package main

import "fmt"

// enumerated types- we implement enum using const, esa koi inbuilt enum nahi hai

//type MyType string // custom type bna skte golang mein ese, with help of type and const, we amke enums in golang

//we inmplement enums using const + type

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received  OrderStatus = "received"
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

// func changeOrderStatus(status string) {
// 	fmt.Println("changing order status to", status)
// }

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Confirmed) //here we not pas commas mtlb ki string, const/enums pas kreneg without commas
}
