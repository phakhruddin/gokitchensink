package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IndexFunc(s string, f func(c int) bool) int

func main() {
	s := "Hello! Let's run Go lang"
	//finding index of first space from s string
	fmt.Print(strings.IndexFunc(s, unicode.IsSpace))
}
