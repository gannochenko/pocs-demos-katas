package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

var leakySlice = make([][]byte, 0) // Memory keeps growing

func leakMemory() {
	for {
		// Allocate 1MB of memory in each iteration and store it
		leakySlice = append(leakySlice, make([]byte, 1024*1024))
		time.Sleep(100 * time.Millisecond) // Prevent CPU overuse
	}
}

func logMemoryUsage() {
	var memStats runtime.MemStats
	for {
		runtime.ReadMemStats(&memStats)
		fmt.Printf("HeapAlloc: %d KB, NumGoroutine: %d\n",
			memStats.HeapAlloc/1024, runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}

func main() {
	go leakMemory() // Start memory-leaking function
	go logMemoryUsage()

	// Start pprof server
	fmt.Println("pprof server running on :6060")
	http.ListenAndServe("localhost:6060", nil)
}
