package main

import (
	"fmt"
	"log"
)

func trace(funcName string) func() {
	log.Printf("Entering %s", funcName)
	return func() {
		log.Printf("Leaving %s", funcName)
	}
}

// Implement isPrime and isComposite functions here

func isPrime(num int) bool {
	defer trace("isPrime")()
	// Implement isPrime logic here
	for i := num; i > 0; i-- {
		log.Println("i =", i)
		check1 := num % i
		log.Println(num, "%", i, "=", check1)
		if check1 == 0 && i != num {
			return false
		} else if check1 != 0 && i != num && i != 1 {
			return true
		}
	}
	return isComposite(num)
}

func isComposite(num int) bool {
	defer trace("isComposite")()
	// Implement isComposite logic here
	log.Println("check", num, "if isComposite ")
	if num%2 == 0 {
		log.Println("isComposite check TRUE for ", num)
		return true
	}
	return isPrime(num)
}

func main() {
	fmt.Printf("%d is a prime number: %t\n\n", 7, isPrime(7))
	fmt.Printf("%d is a composite number: %t\n\n", 8, isComposite(8))
	fmt.Printf("%d is a prime number: %t\ni\n", 9, isPrime(9))
	fmt.Printf("%d is a prime number: %t\n\n", 10, isPrime(10))
	fmt.Printf("%d is a composite number: %t\n\n", 10, isComposite(10))
}
