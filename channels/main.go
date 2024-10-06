// //its like a pipe from whhch we send data from one side and reive data from other side.
// //ek goroutine se dusri goroutine ke andr data send krte hain toh with the help of channel krte hain, bcz ata atime bhut sare gorutines run ho rhe hote hain
// //it measn the communication between the goroutines we can do with the help of channels

// package main

// import (
// 	"fmt"
// )

// //sending data

// // func processNum(numChan chan int) {
// // 	//goroutine ke andr data receive krte hain
// // 	for num := range numChan {
// // 		fmt.Println("processing number", num)
// // 		time.Sleep(time.Second)
// // 	}

// // }

// // receive data
// // func sum(result chan int, num1 int, num2 int) {
// // 	numResult := num1 + num2
// // 	result <- numResult

// // }

// // go routine synchronizer
// // func task(done chan bool) {
// // 	defer func() {
// // 		done <- true
// // 	}()
// // 	fmt.Println("processing")
// // 	// done <- true
// // }

// // queue sys tem implenting
// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 	}
// }

// func main() {

// 	//done := make(chan bool) //unbuffered channel- iski sending or reciving blocking hai, ye udr use hota hai jidr hme one-by-one dats send recive krna hai like with the help opf blocking
// 	//but if we want to implement something interesting like queue system, we can use here bufferef channels. isme hum limited amount of data send kr skte hain without blocking

// 	emailChan := make(chan string, 5)
// 	done := make(chan bool)

// 	for i := 0; i < 5; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i)
// 	} // realtime ye emails database email query krkr get krenge hum

// 	fmt.Println("done sending..")

// 	// emailChan <- "1@ex.com"
// 	// emailChan <- "2@ex.com"

// 	// fmt.Println(<-emailChan)
// 	// fmt.Println(<-emailChan)
// 	close(emailChan) //closing the channel is important
// 	<-done

// 	//go task(done)

// 	//<-done //blocking

// 	// result := make(chan int)

// 	// go sum(result, 4, 5)
// 	// res := <-result //blocking
// 	// fmt.Println(res)

// 	//creat cahnnel syntax
// 	// messageChan := make(chan string)

// 	// messageChan <- "ping" //blocking
// 	// msg := <-messageChan
// 	// fmt.Println(msg)

// 	//deadlock mtlb jb os mein multiple processes chl rhe hain concurrwntly, ye processes kisi resource ko hold krlete hain

// 	// numChan := make(chan int)
// 	// go processNum(numChan)
// 	// for {
// 	// 	numChan <- rand.Intn(100)
// 	// }

// 	// numChain <- 5
// }

// //we can us e queue system in our sppliocstion with thehelp of channel
// // we can do synchronization of threads with the helpp of channels
// //sending and receiving blocking hote hain.
// // if there is single goroutine toh channels use kr skte for synchromiztaion of threads and id there is multiple multiple routines toh  waitgroups use kro, vo jyada flexibiliry provide krte hain.

// //

// queue system implemenataion
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	emailChan := make(chan string, 100) //buffered channel

// 	done := make(chan bool)

// 	go emailSender(emailChan, done)

// 	for i := 1; i <= 5; i++ {
// 		emailChan <- fmt.Sprintf("%d@ex.com", i)

// 	}

// 	fmt.Println("done sending")
// 	// emailChan <- "1@ex.com"
// 	// emailChan <- "2@ex.com"

// 	// fmt.Println(<-emailChan)
// 	// fmt.Println(<-emailChan)
// 	close(emailChan)
// 	<-done

// }

// if we want to listen on multiple channels
package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	// inline go routine annyonous func
	//closure use kr rhe idr
	go func() {
		chan1 <- 10

	}()

	go func() {
		chan2 <- "poo"
	}()

	//recive data
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("recived data from chan1", chan1Val)
		case chan2val := <-chan2:
			fmt.Println("recived data from chan2", chan2val)
		}

	}
}

//we can also make chnnel jispr hum sirf channel recive hi kr skte hain or send hi kr skte. generlly ek channel pr hum doo kr skte. but hum ek specific chij krna chahte vo bhi kr skte hain.

// recive only channel
func emailSender(emailChan <-chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

//send only channel
// func emailSender(emailChan <-chan string, done chan->(arrow ki direction change krdi bs) bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }
