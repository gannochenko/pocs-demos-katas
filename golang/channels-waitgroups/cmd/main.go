package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	numCPUs := runtime.NumCPU()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	signalChannel := GetSignalChan()
	go func() {
		// reading from the channel. this is a blocking operation
		<-signalChannel

		fmt.Println("Terminating...")
		cancelCtx()
	}()

	wg := sync.WaitGroup{}
	wg.Add(numCPUs) // add numCPUs locks to the wait group

	for i := 0; i < numCPUs; i++ {
		// fly, my friend
		go ChuckNorris(ctx, &wg, fmt.Sprintf("Chuck #%d", i), int32(i*10))
	}

	go func() {
		// after 5 seconds the program decides to exit
		time.Sleep(5 * time.Second)
		cancelCtx()
	}()

	// wait until all Chucks gracefully shut down
	wg.Wait()
}

func ChuckNorris(ctx context.Context, wg *sync.WaitGroup, id string, increment int32) {
	counter := int32(0)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s has left the building\n", id)
			wg.Done() // release 1 lock from the wait group
			return
		default:
		}

		// Sleep for 2 seconds
		time.Sleep(2 * time.Second)

		counter += increment
		fmt.Printf("%s counts: %d\n", id, counter)
	}
}

func GetSignalChan() chan os.Signal {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	return signalChannel
}
