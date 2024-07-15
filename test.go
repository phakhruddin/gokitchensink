package main

import "fmt"

func main() {
	type foo struct {
		msg string
		cnt int
	}
	f := foo{"hello", 0}
	for f.cnt < 3 {
		fmt.Println(f.msg)
		f.cnt++
	}
}
