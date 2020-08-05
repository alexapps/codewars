package main

import (
	"fmt"
	"sync"
)

type A struct {
	sync.Mutex
	MyMap map[int]struct{}
}

func main() {
	fmt.Println("Test")
	// b := &A{}
	if getA() == nil {
		fmt.Println("Equal")
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	// var m sync.Mutex
	a := &A{mu, new(map[int]struct{})}
	for i := 0; i < 10; i++ {
		// a.Lock()
		a.I = i
		// mu.Lock()
		go Some(a, &wg)
		wg.Add(1)
		wg.Wait()
		// a.Unlock()
	}

}

func Some(conf *A, wg *sync.WaitGroup) {
	// conf.Lock()
	fmt.Println(conf.I)
	// conf.Unlock()
	// time.Sleep(time.Millisecond)
	defer func() {

		wg.Done()
	}()

}

func getA() *A {
	return &A{}
}
