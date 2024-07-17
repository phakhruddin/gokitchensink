package main

import (
	"fmt"
	"strconv"
	"strings"
)

func IdentifyPrefixPostfix(userID, email string) bool {
	return strings.HasPrefix(userID, "UID-") || strings.HasSuffix(email, "io")
}

func ContainsEducative(email string) bool {
	return strings.Contains(email, "educative")
}

func MaskUserName(email string) string {
	// Implement this function
	//return strings.Replace("XXXX", email, "", 0)
	x := strings.Split(email, "@")
}

func IndexOfAtSymbol(email string) int {
	// Implement this function
	return strings.Index(email, "@")
}

func TrimAndSplitUserID(userID string) string {
	// Implement this function
	//return strings.TrimSpace(userID)
	return strings.TrimPrefix(userID, "UID-")
}

func ConvertStringToInt(str string) int {
	// Implement this function
	var e error
	y, e := strconv.Atoi(str)
	if e == nil {
		return y
	}
	return y
}

func main() {
	// Test your functions here
	fmt.Println(IdentifyPrefixPostfix("UID-0123", "evangeline@educative.io")) // true
	fmt.Println(ContainsEducative("evangeline@educative.io"))                 // true
	fmt.Println(MaskUserName("evangeline@educative.io"))                      // e******e@educative.io
	fmt.Println(IndexOfAtSymbol("evangeline@educative.io"))                   // 10
	fmt.Println(TrimAndSplitUserID("UID-0123"))                               // 0123
	fmt.Println(ConvertStringToInt("123"))                                    // 123
}
