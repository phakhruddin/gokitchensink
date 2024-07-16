package main

import "fmt"

// Global variable
var total int

// FIX START
const (
	Factor1 = 2
	Factor2 = 3
)

// FIX END

func main() {
	// local variable
	var sum int

	{
		// Block level local variable
		count := 5
		sum = calculateSum(&count, 2)
		fmt.Println("calculateSum Sum:", sum)
		//
		sum = calculateSumPlain(count, 2)
		fmt.Println("calculateSumPlain Sum:", sum)
	}
	total = sum
	{
		count := 10
		sum = calculateSum(&count, 3)
		fmt.Println("Sum:", sum)
	}
	total := total + sum
	fmt.Println("Total:", total)

	//FIX START
	// Initialize variables within the smallest possible scope
	sum1 := calculateSumF(5, Factor1)
	fmt.Println("Sum1:", sum1)

	sum2 := calculateSumF(10, Factor2)
	fmt.Println("Sum2:", sum2)

	// Calculate total directly without redeclaration
	total1 := sum1 + sum2
	fmt.Println("Total1:", total1)
	// FIX END

}

func calculateSum(count *int, factor int) int {
	// Accessing local variable
	sum := *count * factor
	return sum
}

func calculateSumPlain(count, factor int) int {
	// Accessing local variable
	sum := count * factor
	return sum
}

func calculateSumF(count, factor int) int {
	// Accessing local variable
	return count * factor
}
