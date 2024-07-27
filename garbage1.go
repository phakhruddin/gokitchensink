package main

import (
	"fmt"
)

var arr []byte = []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}

func main() {
	fmt.Println("arr []byte, ", arr)
	arru := make([]byte, len(arr)) // this will contain consecutive non-repeating items
	ixu := 0                       // index in arru
	tmp := byte(0)
	for _, val := range arr {
		fmt.Println("val, ", val, " tmp, ", tmp)
		if val != tmp {
			arru[ixu] = val
			//fmt.Printf("%c ", arru[ixu])
			fmt.Printf("%c \n", arru[ixu])
			ixu++
		}
		tmp = val
	}
}
