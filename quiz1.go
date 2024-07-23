package main

import "fmt"

func DoSomething1(a *int) {
	var b *int
	b = a
	return b
}

func DoSomething2(a int) {
	var c int
	c = &a
	return c
}

func main() {
	fmt.Println(DoSomething1(1))
	fmt.Println(DoSomething2(1))
}
