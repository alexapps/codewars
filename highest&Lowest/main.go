package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	// Code here or
	tmp := ""
	numbSlice := []int{}
	someStr := strings.Split(in, "")
	fmt.Println(someStr)
	fmt.Println(someStr[2])
	for _, v := range in {

		if v == 32 {
			continue
		}
		number, err := strconv.Atoi(string(v))
		if err != nil {
			return tmp
		}
		numbSlice = append(numbSlice, number)
	}
	sort.Ints(numbSlice)
	// Prepare output

	return fmt.Sprintf("%d %d", numbSlice[0], numbSlice[len(numbSlice)-1])
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}
