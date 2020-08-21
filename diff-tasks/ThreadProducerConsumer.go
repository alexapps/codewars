package main

/*
 http Server and run in go
*/

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	// Producer
	go func(ch chan string) {
		for i := 0; i < 5; i++ {

			ch <- "Some " + strconv.Itoa(i)
		}
		close(ch)
		wg.Done()
	}(ch)

	wg.Add(1)
	// Consumer
	go func(ch chan string) {
		for {
			for data := range ch {
				fmt.Println("Data is ", data)
			}
		}
	}(ch)
	// wg.Add(1)
	wg.Wait()

}
