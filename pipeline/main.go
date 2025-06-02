package main

import (
	"fmt"
	"sync"
)

func Doubler(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		out <- val * 2
	}
	close(out)
}

func Halver(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		out <- val / 2
	}
	close(out)
}

func Printer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		fmt.Println("result value:", val)
	}
}

func Generator(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func main() {
	var wg sync.WaitGroup

	genCh := make(chan int)
	doubleCh := make(chan int)
	halveCh := make(chan int)

	wg.Add(4)
	go Generator(genCh, &wg)
	go Doubler(genCh, doubleCh, &wg)
	go Halver(doubleCh, halveCh, &wg)
	go Printer(halveCh, &wg)

	wg.Wait()
}
