package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(CountBits(1229))
}

func CountBits(num uint) int {
	return bits.OnesCount64(uint64(num))
}
