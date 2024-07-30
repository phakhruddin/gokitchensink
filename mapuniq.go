package main

import (
	"fmt"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43,
		"lima": 98}
)

func main() {
	invMap := make(map[int][]string, len(barVal)) // interchanging types of keys and values
	for k, v := range barVal {
		var a [1]string           // making an array to hold values for a key in invMap
		a[0] = k                  // placing key into array before passing directly into invMap
		_, isPresent := invMap[v] // checking existence
		if isPresent {            // if already present
			fmt.Println("invMap[v]:", invMap[v])
			invMap[v] = append(invMap[v], ",") // add comma before adding another value
		}
		invMap[v] = append(invMap[v], a[0]) // key becomes value and value becomes key
	}
	fmt.Println("inverted:")
	for k, v := range invMap {
		fmt.Printf("key: %v, value: %v / ", k, v)
	}
}
