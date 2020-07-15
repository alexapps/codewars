package main

import "fmt"

func main() {
	fmt.Println("Test")
	a := []int64{1, 2, 3}
	b := []int64{67}
	c := []int64{67, 78, 78}
	MostPopularNumber(a, 3)
	MostPopularNumber(b, 1)
	MostPopularNumber(c, 3)
}

func MostPopularNumber(sourceSlice []int64, sourceLngth int) int64 {
	if sourceLngth == 1 {
		return sourceSlice[0]
	}
	// use for store the result
	storeMap := make(map[int64]int, 0)
	for _, value := range sourceSlice {
		v, ok := storeMap[value]
		if ok {
			storeMap[value] = v + 1
		} else {
			storeMap[value] = 1
		}
	}
	fmt.Println(storeMap)
	return 0
}
