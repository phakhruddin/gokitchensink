package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	//	parameterValueParse := flag.Parse()
	parameterValue := flag.Args()
	parameterValueParse := flag.Parsed()

	//	log.Printf("parameter is: %s\n", parameterName)
	fmt.Printf("String: %+s\n", parameterValue)
	fmt.Printf("String: %+q\n", parameterValueParse)

}
