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

	// need to explicitly call make(), as "var outputChannel chan int32" will not work
	outputChannel := make(chan int32)

	signalChannel := GetSignalChan()
	go func() {
		// reading from the channel. this is a blocking operation
		<-signalChannel

		fmt.Println("Terminating...")
		cancelCtx()
		close(signalChannel)
	}()

	wg := sync.WaitGroup{}
	wg.Add(numCPUs) // add numCPUs locks to the wait group

	for i := 0; i < numCPUs; i++ {
		// fly, my friend
		go ChuckNorris(ctx, &wg, outputChannel, fmt.Sprintf("Chuck #%d", i), int32(i*10))
	}

	for {
		msg, open := <-outputChannel
		if !open {
			break
		}
		fmt.Println("Received:", msg)
	}

	// wait until all Chucks gracefully shut down
	wg.Wait()
}

func ChuckNorris(ctx context.Context, wg *sync.WaitGroup, outputChannel chan int32, id string, increment int32) {
	counter := int32(0)
	sent := true
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

		if sent {
			counter += increment
			fmt.Printf("%s counts: %d\n", id, counter)
		}

		select {
		case outputChannel <- counter:
			sent = true
		default:
			sent = false
		}
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
