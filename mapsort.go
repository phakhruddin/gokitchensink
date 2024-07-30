package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23, "delta": 87,
		"echo": 56, "foxtrot": 12, "golf": 34, "hotel": 16, "indio": 87, "juliet": 65, "kilo": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key: %v, value: %v / ", k, v) // read random keys
	}
	keys := make([]string, len(barVal)) // storing all keys in separate slice
	i := 0
	for k := range barVal {
		fmt.Println(k)
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	fmt.Println(barVal["bravo"])
	sort.Strings(keys) // sorting the keys slice
	fmt.Println(keys)
	fmt.Println(barVal["bravo"])

	fmt.Println()
	fmt.Println("\nsorted:")
	for _, k := range keys {
		fmt.Printf("key: %v, value: %v / ", k, barVal[k]) // reading key from keys and value from barVal
	}

	drinks := map[string]string{
		"beer":   "bière",
		"wine":   "vin",
		"water":  "eau",
		"coffee": "café",
		"thea":   "thé"}
	sdrinks := make([]string, len(drinks))
	ix := 0
	for eng := range drinks {
		sdrinks[ix] = eng
		ix++
	}
	sort.Strings(sdrinks)
	for _, eng := range sdrinks {
		fmt.Println(eng, ":", drinks[eng])
	}
}
