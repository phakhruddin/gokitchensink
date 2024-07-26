package main

import "fmt"

func main() {
	v := make([]int, 10, 50)
	fmt.Println(v)
	p := new([]int)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

}
