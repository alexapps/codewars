package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(IsAlmostPalindrome("qwwq"))

}

func IsAlmostPalindrome(word string) bool {
	wordAsByte := []byte(word)
	wordLength := len(wordAsByte)
	// Check first is word could be an p.
	if wordLength%2 != 0 {
		return false
	}
	// get the first part
	part1 := wordAsByte[0 : wordLength/2]
	part2 := wordAsByte[wordLength/2 : wordLength]
	// Simple comparission
	// Reverse the part2
	reverseSlice(part2)
	if reflect.DeepEqual(part1, part2) {
		return true
	} else {
		return false
	}
	return true
}

func reverseSlice(data interface{}) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice {
		fmt.Println("Should be slice")
	}
	valueLen := value.Len()
	for i := 0; i <= int((valueLen-1)/2); i++ {
		reverseIndex := valueLen - 1 - i
		tmp := value.Index(reverseIndex).Interface()
		value.Index(reverseIndex).Set(value.Index(i))
		value.Index(i).Set(reflect.ValueOf(tmp))
	}
}
