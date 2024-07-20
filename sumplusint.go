package main

import "fmt"

func sumInts(list ...int) (sum int) {
	for _, v := range list {
		fmt.Println(v)
		sum = sum + v
	}
	return
}

func main() {
	fmt.Println(sumInts(5, -2, 0, 9, 10, 11))
}
