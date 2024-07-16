package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

func main() {
	temp := toFahrenheit(100)
	fmt.Println("toFahrenheit(t)", temp)
}

// Function to convert celsius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit
	temp = (Fahrenheit(t) * 9 / 5) + 32
	return temp
}
