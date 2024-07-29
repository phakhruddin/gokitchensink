package main

import "fmt"

var Days = map[int]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday"}
var Dayname = make([]string, 10)
var isPresent bool
var value string

func findDay(n int) string {
	value, isPresent = Days[n]
	if isPresent {
		return value
	} else {
		return "Wrong key"
	}
}

func main() {
	Dayname := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for i := 0; i < len(Dayname); i++ {
		//fmt.Println(i, Dayname[i])
		Days[i+1] = Dayname[i]
	}
	//fmt.Println(Days)
	fmt.Println(findDay(7))
}
