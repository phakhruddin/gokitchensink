package main

import "fmt"

var initvP *uint64

func Factorial(n uint64) (fac uint64) {
	if n <= 1 { //base case
		return 1
	}
	fac = n * Factorial(n-1) // recursive case
	return
}

func FactorialMy(n uint64) (fac uint64) {
	//var initv uint64
	//initv = 1
	switch n {
	case 0, 1:
		return 1
	case 2:
		return 2
	default:
		initvP = &n
		//fmt.Printf("The value at memory location %p is %d\n", initvP, *initvP)

		for i := n; i > 2; i-- {
			fac = *initvP * (i - 1)
			//fmt.Printf("n is %d and fac is %v at %d\n", n, fac, i)
			initvP = &fac
			//fmt.Printf("The new value at memory location %p is %d\n", initvP, *initvP)
		}
		return fac
	}
}

func main() {
	var num uint64
	num = 16
	for i := num; i > uint64(0); i-- {
		fmt.Printf("Factorial(%x) is %v \n", i, Factorial(i))
		fmt.Printf("FactorialMy(%x) is %v \n", i, FactorialMy(i))
	}

}
