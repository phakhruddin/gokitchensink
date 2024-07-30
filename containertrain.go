package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
