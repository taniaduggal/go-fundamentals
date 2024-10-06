// package main

// import "fmt"

// // func main() {
// // 	fmt.Println("welcome to functions")
// // 	greeter()

// // 	result := adder(3, 5)
// // 	fmt.Println("Result is:", result)

// // 	proRes, myMess := proAdder(2, 4, 5, 6, 7, 8)
// // 	fmt.Println("Result is:", proRes, myMess)
// // }

// // //variadic functions: ...
// // func proAdder(values ...int) (int, string) {
// // 	total := 0

// // 	for _, val := range values {
// // 		total += val
// // 	}
// // 	return total, "hello doing for demo purposes"
// // }

// // func adder(valOne int, valTwo int) int {
// // 	return valOne + valTwo
// // }

// // // even anonymous functions = lambda functions exist without name, u can say its a self invoked funcs like this
// // // func () {

// // // } ()

// // func greeter() {
// // 	fmt.Println("hyy world")
// // }

// //first class citizen function hote go mein - it means we can return ,pass it as an argument in another function. assign to a variable

// // func processIt(fn func(a int) int) {
// // 	fn(1)
// // }

// // func processIt() func(a int) int { //i returns the func
// // 	return func(a int) int {
// // 		return 4
// // 	}
// // }

// // func main() {
// // 	// fn := func(a int) int { //anonymous func
// // 	// 	return 2
// // 	// }
// // 	fn1 := processIt()

// // }

// // Closures are a special case of anonymous functions. Closures are anonymous functions which access the variables defined outside the body of the function.
// //Anonymous function accessing the variable defined outside body.
// // func main() {
// // 	l := 20
// // 	b := 30

// // 	func() {
// // 		var area int
// // 		area = l * b
// // 		fmt.Println(area)
// // 	}()
// // }

// // func main() {
// // 	for i := 10.0; i < 100; i += 10.0 {
// // 		rad := func() float64 {
// // 			return i * 39.370
// // 		}()
// // 		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
// // 	}
// // }

// func sum(x, y int) int {
// 	return x + y
// }
// func partialSum(x int) func(int) int {
// 	return func(y int) int {
// 		return sum(x, y)
// 	}
// }
// func main() {
// 	partial := partialSum(3)
// 	fmt.Println(partial(7))
// }
