package main

import (
	"bytes"
	"fmt"
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
	fmt.Println("Test", isBalanced("[jhghjg]"))

}

var (
	opened = []byte{40, 123, 91}
	close  = []byte{41, 125, 93}
)

func isBalanced(str string) bool {
	retValue := false
	strAsByte := []byte(str)
	start, end := initTheCurlyPosition()

	// fill all
	for i := 0; i < len(str); i++ {
		cuurentStart := false
		currendEnd := false
		if bytes.Contains(opened, byte(strAsByte[i])) {
			cuurentStart = true
			setTheCurlyPosition(start, strAsByte[i], i)
		}
		if bytes.Contains(close, byte(strAsByte[i])) {
			currendEnd = true
			setTheCurlyPosition(end, strAsByte[i], i)
		}
	}
	// TODO go trough the slices and check....

	return retValue
}

// The key is ascci code of braket, pos - new
func setTheCurlyPosition(inMap map[int][]int, key, pos int) {
	inMap[key] = append(inMap[key], pos)
}

func initTheCurlyPosition() (map[int][]int, map[int][]int) {
	start := make(map[int][]int, 3)
	end := make(map[int][]int, 3)
	start[40] = []int{}
	start[123] = []int{}
	start[91] = []int{}
	end[41] = []int{}
	end[125] = []int{}
	end[93] = []int{}

}
