package main

import "fmt"

var slice1 = make([]string, 10)

// slice is the slice to which another slice will be added
// insertion is the slice that needs to be inserted.
// index is the position for insertion
/*func insertSlice(slice, insertion []string, index int) []string {
	fmt.Println("slice = ", slice, "insertion = ", insertion, "index = ", index)
	slice1 = append(insertion, slice[index:]...)
	slice1 = append(slice[:index], slice1...)
	return slice1
}*/

func main() {
	fmt.Println(insertSlice([]string{"M", "N", "O", "P", "Q", "R"}, []string{"A", "B", "C"}, 2))
}

// also another option
func insertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	fmt.Println("at :=, ", at, " result, ", result)
	at += copy(result[at:], insertion)
	fmt.Println("at +=, ", at, " result, ", result)
	copy(result[at:], slice[index:])
	return result
}
