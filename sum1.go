package main

import "fmt"

// Constants for factor values
const (
	Factor1 = 2
	Factor2 = 3
)

func main() {
	// Initialize variables within the smallest possible scope
	sum1 := calculateSum(5, Factor1)
	fmt.Println("Sum:", sum1)

	sum2 := calculateSum(10, Factor2)
	fmt.Println("Sum:", sum2)

	// Calculate total directly without redeclaration
	total := sum1 + sum2
	fmt.Println("Total:", total)
}

func calculateSum(count, factor int) int {
	// No need to modify count, just use it in the expression
	return count * factor
}
