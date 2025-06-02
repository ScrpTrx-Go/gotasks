package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	jobsChan := make(chan int)
	resChan := make(chan int, 50)
	go func() {
		for i := 0; i < 50; i++ {
			jobsChan <- i
		}
		close(jobsChan)
	}()

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case val, ok := <-jobsChan:
					if !ok {
						return
					}
					fmt.Printf("Gourutine %d works with value %d\n", id, val)
					resChan <- val
				}
			}
		}(i)
	}
	wg.Wait()
	close(resChan)
	for values := range resChan {
		fmt.Println(values)
	}
}
