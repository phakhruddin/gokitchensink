package main

import "fmt"

func fibarray(term int) []int {
	nil := make([]int, term, 10)
	nil[0] = 0
	//var slice []int = arr{1: 5}
	for i := 1; i < len(nil); i++ {
		if i > 2 {
			nil[i] = nil[i-1] + nil[i-2]
		} else {
			nil[i] = 1
		}
	}
	return nil
}

func main() {
	fmt.Println(fibarray(5))
}
