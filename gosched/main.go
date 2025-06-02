package main

import (
	"fmt"
	"runtime"
	"time"
)

func goroutineWithGosched(name string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s: %d\n", name, i)
		if i == 50 {
			runtime.Gosched()
		}
	}
}

func goroutineWithoutGosched(name string) {
	for i := 100; i < 150; i++ {
		fmt.Printf("%s: %d\n", name, i)
		if i == 110 {
			runtime.Gosched()
		}
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go goroutineWithGosched("1")
	go goroutineWithoutGosched("2")

	time.Sleep(1 * time.Second)
}
