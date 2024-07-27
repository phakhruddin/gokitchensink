package main

import (
	"fmt"
)

func bubbleSort(sl []int) []int {
	fmt.Println("sl, ", sl)
	for h := 1; h < len(sl); h++ {
		for i := 0; i < len(sl)-h; i++ {
			fmt.Println("at h:", h, " at i:", i, ",sl[i]:", sl[i], ",vs. sl[i+1]:", sl[i+1], ", sl:", sl)
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
				fmt.Println("at h:", h, " at i:", i, ",sl[i]:", sl[i], ",vs. sl[i+1]:", sl[i+1], ", sl:", sl)
			}
		}
	}
	return sl
}

func main() {
	fmt.Println(bubbleSort([]int{4, 5, 2, 1, 3}))

}
