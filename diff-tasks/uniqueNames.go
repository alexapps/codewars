package main

import "fmt"

func uniqueNames(a, b []string) []string {
	var result []string

	tmpMap := make(map[string]struct{})
	for _, v := range a {
		tmpMap[v] = struct{}{}
	}
	for _, v := range b {
		tmpMap[v] = struct{}{}
	}
	i := 0
	result = make([]string, len(tmpMap))
	for k := range tmpMap {
		result[i] = k
		i++
	}
	return result
}

func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
		[]string{"Ava", "Emma", "Olivia"},
		[]string{"Olivia", "Sophia", "Emma"}))
}
