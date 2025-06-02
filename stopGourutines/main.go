package main

import (
	"fmt"
	"sync"
)

func main() {
	var wgInit sync.WaitGroup
	var wgDone sync.WaitGroup
	var mu sync.RWMutex
	stopGourutines := make(map[int]chan struct{})
	for i := 0; i < 5; i++ {
		wgDone.Add(1)
		wgInit.Add(1)
		go func(id int) {
			defer wgDone.Done()
			stopChan := make(chan struct{}, 1)
			mu.Lock()
			stopGourutines[id] = stopChan
			mu.Unlock()
			fmt.Printf("gourutine %d started\n", id)
			wgInit.Done()
			select {
			case <-stopChan:
				fmt.Printf("gourutine %d finished\n", id)
				return
			}
		}(i)
	}
	wgInit.Wait()
	go func() {
		for i := 0; i < 5; i++ {
			mu.RLock()
			stopChan, ok := stopGourutines[i]
			mu.RUnlock()
			if ok {
				fmt.Printf("send stop signal to %d gourutine\n", i)
				close(stopChan)
			}
		}
	}()
	wgDone.Wait()
}
