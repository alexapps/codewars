package main

import (
	"fmt"
	"sort"
)

/*
Complete the method which returns the number which is most frequent
in the given input array. If there is a tie for most frequent number,
return the largest number among them.

Note: no empty arrays will be given.

Examples
[12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
[12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
[12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3
*/

func HighestRank(nums []int) int {
	pairs := make(map[int]int)
	for _, v := range nums {
		pairs[v]++
	}
	// fmt.Println("-> map ", pairs)

	type kv struct {
		Key   int
		Value int
	}
	var ss []kv
	for k, v := range pairs {
		ss = append(ss, kv{
			Key:   k,
			Value: v,
		})
	}

	// sort
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss[0].Key
}

func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12, 10}))
	fmt.Println(HighestRank([]int{12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10}))
}
