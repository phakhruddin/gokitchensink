package main

import "fmt"

func main() {
	// implement your loop here
	var arr [15]int
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}
