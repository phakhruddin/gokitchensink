package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))   // 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))   // 18 is odd: is false
}

func even(nr int) bool {
	fmt.Println("even(", nr, ")")
	if nr == 0 {
		return true
	}
	return odd(AbsVal(nr) - 1) // even calls odd
}

func odd(nr int) bool {
	fmt.Println("odd(", nr, ")")
	if nr == 0 {
		return false
	}
	return even(AbsVal(nr) - 1) // odd calls even
}

func AbsVal(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
