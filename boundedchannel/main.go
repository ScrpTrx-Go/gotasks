package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	boundedCh := make(chan int, 10)

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 20; i++ {
			boundedCh <- i
			fmt.Println("zapisano:", i)
		}
		close(boundedCh)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		for val := range boundedCh {
			fmt.Println("pro4itano:", val)
		}
	}()
	wg.Wait()
}
