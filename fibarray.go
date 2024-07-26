package main

import "fmt"

var fib [10]int64

func fibs() [10]int64 {
	for i := int64(0); i < 10; i++ {
		if i >= 2 {
			fib[i] = fib[i-1] + fib[i-2]
		} else {
			fib[i] = 1
		}

	}
	return fib
}

func main() {
	fmt.Println(fibs())
}
