package main

import (
	"fmt"
	"sync"
	"time"
)

func safeGo(panicChan chan<- any, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			panicChan <- r
		}
	}()
	fn()
}

func main() {
	var wg sync.WaitGroup

	panicChan := make(chan any, 1)

	wg.Add(1)
	go safeGo(panicChan, func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		// panic("something wrong")
		fmt.Println("test")
	})
	go func() {
		if result, ok := <-panicChan; ok {
			panic(result)
		}
	}()
	wg.Wait()
}
