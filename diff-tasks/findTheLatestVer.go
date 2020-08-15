package main

/**
* 1. Create random valued 100 elements of struct (name, version, isActive). E.g {"Some", "1.02", false}
* 2. Found the max ver of active and print it out
* 3. After step 2 found the non active with the first value that bigger than max active
 */

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// random generator
var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

// A -
type A struct {
	Name     string
	Version  string
	IsActive bool
}

func main() {
	active := make(map[string]A)
	nonActive := make(map[string]A)

	for i := 0; i < 100; i++ {
		// for name
		name := RandStringRunes(10)
		// for version
		verStr := fmt.Sprintf("%.2f", randFloats(1.0, 9.9)) // fix
		// is active
		tempBool := r.Intn(2) != 0
		isActive := tempBool

		if isActive {
			active[verStr] = A{name, verStr, isActive}
		} else {
			nonActive[verStr] = A{name, verStr, isActive}
		}
	}
	// // Test
	fmt.Println("**********")
	fmt.Println("active ", active)
	fmt.Println("**********")
	fmt.Println("non-active ", nonActive)
	fmt.Println("**********")

	// Get the max ver active element
	activeKeys := getSortedKeys(active)
	maxVer := ""
	if activeKeys != nil && len(activeKeys) > 0 {
		maxVer = fmt.Sprintf("%.2f", activeKeys[len(activeKeys)-1])
	}
	// Printout the task 1
	maxVerF, err := strconv.ParseFloat(maxVer, 64)
	if err != nil {
		// todo
	}
	maxActiveValue, _ := active[maxVer]
	fmt.Printf("Active max ver %s is %v\n", maxVer, maxActiveValue)

	// Step2
	nonActiveKeys := getSortedKeys(nonActive)
	nearKey := ""
	for _, nonActiveKey := range nonActiveKeys {
		if nonActiveKey < maxVerF {
			continue
		} else if nonActiveKey > maxVerF {
			nearKey = fmt.Sprintf("%.2f", nonActiveKey)
			break
		}
	}
	nonActiveValue, _ := nonActive[nearKey]
	fmt.Printf("Next to the Active max ver a non active version %s is %v\n", nearKey, nonActiveValue)

}

func getSortedKeys(m map[string]A) (keys []float64) {
	keys = make([]float64, 0, len(m))
	for k := range m {
		s, err := strconv.ParseFloat(k, 64)
		if err != nil {
			// todo
			continue
		}
		keys = append(keys, s)
	}
	// sort
	sort.Float64s(keys)
	return
}

// tools

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randFloats(min, max float64) (res float64) {
	res = min + rand.Float64()*(max-min)
	return
}
