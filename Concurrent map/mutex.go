package main

import (
	"fmt"
	"sync"
	"time"
)

func (i intMap) PutMutex(key string, value int, mu *sync.RWMutex) {
	mu.Lock()
	i[key] = value
	mu.Unlock()
}

func (i intMap) GetMutex(key string, mu *sync.RWMutex) (value int) {
	mu.RLock()
	defer mu.RUnlock()
	for k, v := range i {
		if k == key {
			return v
		}
	}

	return 0
}

func MutexFunc() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	i := make(intMap)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		i.PutMutex("stroka", 1, &mu)
		time.Sleep(100 * time.Millisecond)
		i.PutMutex("stroka1", 2, &mu)
		i.PutMutex("stroka2", 3, &mu)
		i.PutMutex("stroka3", 4, &mu)
		i.PutMutex("stroka4", 5, &mu)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		i.PutMutex("stroka5", 1, &mu)
		i.PutMutex("stroka6", 2, &mu)
		i.PutMutex("stroka7", 3, &mu)
		i.PutMutex("stroka8", 4, &mu)
		i.PutMutex("stroka9", 5, &mu)
		time.Sleep(100 * time.Millisecond)
	}()

	go func() {
		wg.Wait()
		res := i.GetMutex("stroka", &mu)
		res1 := i.GetMutex("stroka1", &mu)
		res2 := i.GetMutex("stroka2", &mu)
		fmt.Println(res, res1, res2)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(i)

}
