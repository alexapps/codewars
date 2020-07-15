package whichAreIn

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Hey")
	a1 := []string{"arp", "live", "strong"}
	a2 := []string{"lively", "alive", "harp", "sharp", "armstrong"}
	_ = InArray(a1, a2)
}

func InArray(array1 []string, array2 []string) []string {
	// return array
	retVal := make([]string, 0)
	// like a set
	set := make(map[string]struct{}, 0)
	// go through the input arrays
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			// Countains should be faster than regExp MustCompile+MatchString
			if strings.Contains(v2, v1) {
				set[v1] = struct{}{}
			}
		}
	}
	fmt.Println(getSortedKeys(set))

	return retVal
}

func getSortedKeys(mymap map[string]struct{}) []string {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
