package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	bufChan := make(chan int, 5)
	altChan := make(chan int, 5)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 11; i++ {
			select {
			case bufChan <- i:
				fmt.Println("записано", i, "len=", len(bufChan))
			default:
				altChan <- i
				fmt.Println("в altChan записан:", i, "len:", len(altChan))
			}
		}
		close(bufChan)
		close(altChan)
	}()

	for val := range bufChan {
		fmt.Println("bufch:", val)
	}

	for val := range altChan {
		fmt.Println("altch:", val)
	}
	wg.Wait()

}
