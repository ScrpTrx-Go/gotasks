package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func router(genNums <-chan int, numChans map[int][]chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for key := range genNums {
		fmt.Printf("work with value %d\n", key)
		channels, ok := numChans[key]
		if !ok {
			fmt.Printf("value %d not exists\n", key)
		}
		for _, channel := range channels {
			fmt.Printf("send to channel %v value %d\n", channel, key)
			channel <- key
		}

	}
	for _, channels := range numChans {
		for _, channel := range channels {
			close(channel)
		}
	}
}

func subscriber1(sub1 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-sub1:
			if !ok {
				return
			}
			fmt.Println("sub1:", val*2)
		}
	}
}

func subscriber2(sub2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-sub2:
			if !ok {
				return
			}
			fmt.Println("sub2:", val+2)
		}
	}
}

func subscriber3(sub3 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-sub3:
			if !ok {
				return
			}
			fmt.Println("sub3:", val-2)
		}
	}
}

func subscriber4(sub4 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-sub4:
			if !ok {
				return
			}
			fmt.Println("sub4:", val/2)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	genNums := make(chan int)
	sub1 := make(chan int)
	sub2 := make(chan int)
	sub3 := make(chan int)
	sub4 := make(chan int)

	numChans := map[int][]chan int{
		1:  {sub1, sub2},
		10: {sub3, sub4},
	}

	go func() {
		for i := 1; i <= 5; i++ {
			val := rand.Intn(10)
			fmt.Printf("send to gennums %d\n", val)
			genNums <- val
		}
		close(genNums)
	}()
	wg.Add(5)
	go router(genNums, numChans, &wg)
	go subscriber1(sub1, &wg)
	go subscriber2(sub2, &wg)
	go subscriber3(sub3, &wg)
	go subscriber4(sub4, &wg)
	wg.Wait()
}
