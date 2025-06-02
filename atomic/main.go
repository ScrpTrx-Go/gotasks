package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type User struct {
	counter int64
}

func (u *User) IncrementValue() {
	atomic.AddInt64(&u.counter, 1)
}

func (u *User) ValueWithAtomic() int64 {
	return atomic.LoadInt64(&u.counter)
}

func noAtomicIncrement(wg *sync.WaitGroup) {
	defer wg.Done()
	u := User{}
	var innerWg sync.WaitGroup

	for i := 0; i < 50; i++ {
		innerWg.Add(1)
		go func() {
			defer innerWg.Done()
			u.counter++
		}()
	}

	innerWg.Wait()
	fmt.Println("increment without atomic", u.counter)
}

func main() {
	var wg sync.WaitGroup
	u := User{}

	for i := 0; i < 50; i++ {
		go func() {
			u.IncrementValue()
		}()
	}

	for u.ValueWithAtomic() < 50 {
	}
	fmt.Println("value with atomic:", u.ValueWithAtomic())
	wg.Add(1)
	noAtomicIncrement(&wg)
	wg.Wait()
}
