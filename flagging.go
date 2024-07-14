package main

import (
	"flag"
	"fmt"
)

var one string

func main() {
	//flag.StringVar(&one, "o", "default", "arg one")
	flag.Parse()
	tail := flag.Args()
	fmt.Printf("Tail: %+q\n", tail)
}
