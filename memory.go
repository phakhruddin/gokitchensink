package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int // Pointer variable
	intP = &i1    // Storing address of i1 in pointer variable
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	// next one
	s := "good bye"
	var p *string = &s
	*p = "ciao" // changing the value at &s

	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}
