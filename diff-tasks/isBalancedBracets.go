package main

import (
	"fmt"
	"regexp"
)

var (
	r = regexp.MustCompile("\\(+|\\)+|\\{+|\\}+|\\[+|\\]+")
)

/*
// (a[0]+b[2c[6]]) {24 + 53} -> true
// f(e(d)) -> true
// [()]{}([]) -> true
// ((b) -> false
// (c] -> false
// {(a[]) -> false
// ([)] -> false
// )( -> false
// -> true
*/

func main() {
	fmt.Println("Test", isBalanced("{()}"))
	// foundSlice := r.FindAllString("l(kj{lj", 10)
	// fmt.Println(len(foundSlice), foundSlice)
}

var (
	opened = []string{"(", "{", "["}
	closed = []string{")", "}", "]"}
)

func isBalanced(str string) bool {
	retValue := true
	// Step 1 define is bracets in the string
	foundSlice := r.FindAllString(str, len(str))
	totalBracets := len(foundSlice)
	fmt.Println("Found bracets ", foundSlice)
	fmt.Println("Found totalBracets ", totalBracets)
	if totalBracets == 0 {
		fmt.Println("No bracets")
		return false
	}
	// Non balanced
	isOdd := totalBracets % 2
	fmt.Println("isOdd, ", isOdd)
	stackToHold := make(chan string, totalBracets)
	defer close(stackToHold)
	if isOdd != 0 {
		return false
	} else {
		for _, v := range foundSlice {
			fmt.Println("v ", v)
			if contains(opened, v) {
				stackToHold <- v
			} else if contains(closed, v) {
				topValue := ""
				select {
				case c := <-stackToHold:
					topValue = c
				default:
					topValue = ""
				}

				if (topValue == "(" && v == ")") ||
					(topValue == "{" && v == "}") ||
					(topValue == "[" && v == "]") {
					fmt.Println("topValue equals")
				} else {
					fmt.Println("topValue not equals")
					return false
				}
			}
		}
		// // check the stack
		// rest := <-stackToHold
		// if "" == rest {
		// 	return true
		// }

	}
	return retValue
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
