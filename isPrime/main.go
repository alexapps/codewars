package main

import (
	"fmt"
)

func IsPrime(n int) bool {

	if n == 0 || n < 0 {
		return false
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(IsPrime(73))

}

