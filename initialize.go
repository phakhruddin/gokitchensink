package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 20, 15, 22, 16}
	fmt.Println(arrAge)
	var arr = [10]int{1, 2, 3}
	fmt.Println(arr)
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	fmt.Println(arrLazy)
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	fmt.Println(arrKeyValue)

}
