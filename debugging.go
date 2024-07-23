package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.Llongfile)
	var where = log.Print

	where1 := func() { // debugging function
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	where1()

	where()
	where1()
	// some code
	a := 2 * 5
	where()
	where1()

	// some more code
	b := a + 100
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	where()
	where1()
}
