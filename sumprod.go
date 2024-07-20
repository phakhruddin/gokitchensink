package main

import "fmt"

func SumProductDiff(i, j int) (s int, p int, d int) { // named version
	s = i + j
	p = i * j
	d = i - j
	return
}

func main() {
	fmt.Println(SumProductDiff(3, 4))
}
