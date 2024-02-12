package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const Interval = time.Second * 3

func main() {
	ctx := context.Background()
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("context cancelled\n")
				return
			case <-time.After(Interval):
				err := doStuff()
				if err != nil {
					cancelCtx()
					return
				}
				fmt.Printf("tick!\n")
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

func doStuff() error {
	return nil
}
