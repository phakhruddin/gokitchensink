package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}
