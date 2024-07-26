package main

import "fmt"

func enlarge(s []int, factor int) []int {
	newsl := make([]int, len(s)*factor)
	copy(newsl, s)
	return newsl
}

func main() {
	sl := []int{1, 2, 3}
	fmt.Println(enlarge(sl, 5))
}
