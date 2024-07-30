package main

import (
	"fmt"
	"regexp"
)

func RemoveHyphensFromPhoneNumbers(text string) string {
	// Write your logic here
	pat := `(\d{3})-(\d{3})-(\d{4})`
	re := regexp.MustCompile(pat)
	str := re.ReplaceAllString(text, "$1$2$3")

	return str
}

func main() {
	text := "John: 123-456-7890 William: 111-222-3333 Steve: 444-555-6666 Thomas: 777-888-9999" // string to search

	updatedText := RemoveHyphensFromPhoneNumbers(text)

	fmt.Println("Text with hyphens removed from phone numbers:")
	fmt.Println(updatedText)
}
