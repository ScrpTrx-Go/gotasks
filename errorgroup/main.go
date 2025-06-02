package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	semChan := make(chan struct{}, 5)

	for i := 0; i < 10; i++ {
		iCopy := i
		runWithSemaphore(ctx, g, semChan, iCopy)
	}

	err := g.Wait()
	if err != nil {
		log.Printf("Ошибка %v", err)
	}
}

func runWithSemaphore(ctx context.Context, g *errgroup.Group, semChan chan struct{}, i int) {

	g.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case semChan <- struct{}{}:
		}
		defer func() { <-semChan }()
		return intworker(ctx, i)
	})
}

func intworker(ctx context.Context, num int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	resnum := num * 2
	if resnum == 4 {
		return fmt.Errorf("num = %d", resnum)
	}
	fmt.Println(resnum)
	return nil
}
