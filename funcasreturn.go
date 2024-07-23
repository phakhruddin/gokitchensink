package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(5))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	var f = Adder2()
	fmt.Print(f(1), " , ")
	fmt.Print(f(20), " , ")
	fmt.Print(f(300))
}

func Add2() func(b int) int { // return a function
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int { // return a function
	return func(b int) int {
		return a + b
	}
}

func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		//	fmt.Println("before: x,", x, " delta,", delta)
		x += delta
		//fmt.Println("after: x,", x, " delta,", delta)
		return x
	}
}
