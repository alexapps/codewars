package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Result: ", ToCamelCase("The_Stealth_Warrior"))
}

func ToCamelCase(s string) string {
	retValue := ""
	// rIsCapital, _ := regexp.Compile("^[A-Z]{1}")
	rIsLower, _ := regexp.Compile("^[a-z]{1}")
	// your code
	// Split string
	for i, word := range regexp.MustCompile("_|-").Split(s, -1) {
		if i != 0 && rIsLower.Match([]byte(word)) {
			retValue = retValue + strings.Title(word)
		} else {
			retValue = retValue + word
		}
	}
	return retValue
}
