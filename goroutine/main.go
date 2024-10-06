// most powerful feature in golangh
// lightwirgth thread, multithreading krni hoti hai, concurrent process krne hote hain

package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task", id)
}

// very important: func main is a goroutine too
func main() {
	for i := 0; i <= 10; i++ {
		// task(i) // it craetes blocking, we want to use power of golang, we want parallley kam kre. we use go keyword
		go task(i) //when we make goroutine toh it is scheduled by a scheduler, scheduler main jati and vo isko schedule krta.
		// go func() {
		// 	fmt.Println(i)
		// }() its a good practice ki hum function reives kra kre and then pass kra kre, its a good practice, see below
		go func(i int) {
			fmt.Println(i)
		}(i)

	}
	time.Sleep(time.Second * 2) // iski helpp se hmne main func ko roka hua hai, generally hum esa nahi krte, uska bhi ek mechanism hai ahmare pas
}

//background workers lik skte, cpu intensive kam kr skte, its a poweful concept
