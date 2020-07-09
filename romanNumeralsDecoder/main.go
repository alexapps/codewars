package main

import "fmt"

func main() {
	fmt.Println(romanToDecimal("MMCDXXI"))
}

func value(value byte) int {
	switch value {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return -1
	}
}

func romanToDecimal(str string) int {
	res := 0
	for i := 0; i < len(str); i++ {
		// TODO
		s1 := value(str[i])
		if i+1 < len(str) {
			s2 := value(str[i+1])
			// Check if the next value is
			if s1 >= s2 {
				res = res + s1
			} else {
				res = res + s2 - s1
				i++
			}
		} else {
			res = res + s1
			i++
		}
	}
	return res
}
