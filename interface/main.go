//properly organise code,scalable

// openClose principle- functions and methods should be open for extensions but close for modifications

// problem - changes, creates problem in testing - iska solution hai interfaces
// for tetsing, we pas the fakemethods, we never pass the actual methods

// interafces - a contract
// helps in developing clean architecture, dependency inversion is a concept jo ye implement kr rha

package main

import "fmt"

//interface has name convention ending with "er"

// interface ke andr method deni hai humko and we are saying jo bhi struct implement krega isko, uske andr ye method honi hi chahiye

type paymenter interface {
	pay(amount float32) //methods
	// pay(amount float32) bool
}

type payment struct {
	gateway paymenter // we not need to give any gateway of concrete implementation
}

func (p payment) pay(amount float32) {
	p.gateway.pay(amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

//jo interface mein hmne method diya hai agr vo method hmne bahr define kra hai and uske siganture bhi match kr rhe hain toh it means interface impliciltly apply ho jayega, no need to use ilke implements keywords like in other langiages

// type fakepayment struct{}

// func (f fakepayment) pay(amount float32) {
// 	fmt.Println("making payment using fake payment for testing purpose")
// }

func main() {
	// 	razorpayPaymentGw := razorpay{}
	// 	newPayment := payment{
	// 		gateway: razorpayPaymentGw,
	// 	}
	// 	newPayment.pay(100.0)
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.pay(100.0)

}
