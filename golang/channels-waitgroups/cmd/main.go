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

type CounterValue struct {
	ChuckID string
	Value   int32
}

func main() {
	numCPUs := runtime.NumCPU()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	// the data channel is used to retrieve the results produced by the other threads
	dataChannel := make(chan *CounterValue)

	signalChannel := GetSignalChan()
	go func() {
		// reading from the channel. this is a blocking operation
		<-signalChannel

		fmt.Println("Terminating...")

		// on terminate we cancel the context and close the data channel, so the main thread could move on
		cancelCtx()
		close(signalChannel)
	}()

	wg := sync.WaitGroup{}
	wg.Add(numCPUs) // add numCPUs locks to the wait group

	for i := 0; i < numCPUs; i++ {
		// fly, my friend
		go ChuckNorris(ctx, &wg, dataChannel, fmt.Sprintf("Chuck #%d", i), int32(i*10))
	}

	// try reading from the channel in an endless cycle. This is a blocking operation,
	// but the main thread doesn't do anything useful anyway
	for {
		msg, open := <-dataChannel
		if !open {
			break
		}
		fmt.Printf("%s counts %d\n", msg.ChuckID, msg.Value)
	}

	// wait until all threads gracefully shut down
	wg.Wait()
}

func ChuckNorris(ctx context.Context, wg *sync.WaitGroup, dataChannel chan *CounterValue, id string, increment int32) {
	counter := int32(0)
	sent := true
	for {
		// check if the context wasn't cancelled
		select {
		case <-ctx.Done():
			fmt.Printf("%s has left the building\n", id)
			wg.Done() // release 1 lock from the wait group
			return
		default:
		}

		// imitate some heavy duty
		time.Sleep(2 * time.Second)

		// do actual work only if the previous one was sent
		if sent {
			counter += increment
		}

		// try sending to the channel
		select {
		case dataChannel <- &CounterValue{
			ChuckID: id,
			Value:   counter,
		}:
			sent = true
		default:
			sent = false
		}
	}
}

// GetSignalChan returns a channel that informs about pressing Ctrl+C
func GetSignalChan() chan os.Signal {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	return signalChannel
}
