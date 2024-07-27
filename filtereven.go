package main

import "fmt"

func Filter(s []int, fn func(int) bool) []int {
	var result []int
	for _, value := range s {
		if even(value) {
			result = append(result, value)
			//fmt.Println("result, ", result)
		}
	}
	return result
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(Filter([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, even))
	fmt.Println(Filter([]int{1, 3, 5, 7}, even))
	fmt.Println(Filter([]int{0, 2, 4, 6, 8, 10}, even))
	fmt.Println(Filter([]int{}, even))
}
