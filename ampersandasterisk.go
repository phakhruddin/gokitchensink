package main

import "fmt"

func main() {
	helloWorld := "helloworld"
	var pointerToHelloWorld *string
	pointerToHelloWorld = &helloWorld
	fmt.Println("Pointer to helloWorld", pointerToHelloWorld)
	fmt.Println("*pointerToHelloWorld", *pointerToHelloWorld) //content of memory address
	fmt.Println("&pointerToHelloWorld", &pointerToHelloWorld) //memory address
	// additional
	var pointerToPointer **string
	//pointerToPointer = &(&helloWorld)
	fmt.Println(pointerToPointer)

}
