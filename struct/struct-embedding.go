package main

import "fmt"

//u have heard of inheritance, composition etc, we wven doing below

type customer struct {
	name  string
	phone string
}

type ordery struct {
	id       string
	amount   float32
	status   string
	customer // struct embedding

}

func main() {
	newCustomer := customer{
		name:  "johbn",
		phone: "123456",
	}
	newOrder1 := ordery{
		id:       "1",
		amount:   30,
		status:   "pending",
		customer: newCustomer,
	}
	newOrder1.customer.name = "robin"
	fmt.Println(newOrder1.customer)
}
