package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func trace(funcName string) func() {
	log.Printf("Entering %s", funcName)
	return func() {
		log.Printf("Leaving %s", funcName)
	}
}

// openFile opens a file and returns a file pointer.
func openFile(filename string) (*os.File, error) {
	defer trace("openFile")()
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil

}

// processFile processes the opened file.
func processFile(file *os.File) error {
	// Do some operations on the file.
	defer trace("processFile")()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}

// closeFile closes the file.
func closeFile(file *os.File) {
	defer trace("closeFile")()
	if file != nil {
		file.Close()
	}
}

func main() {
	fileName := "example.txt"

	file, err := openFile(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer closeFile(file)

	err = processFile(file)
	if err != nil {
		log.Fatalf("Error processing file: %v", err)
	}

	fmt.Println("File processed successfully.")
}
