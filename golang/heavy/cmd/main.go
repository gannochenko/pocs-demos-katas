package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Import pprof
	"time"
)

// Inefficient Fibonacci function (CPU-heavy)
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2) // Exponential complexity: O(2^n)
}

func main() {
	// Start pprof server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Starting heavy computation...")

	start := time.Now()
	result := fibonacci(999) // Change this number to increase/decrease workload
	elapsed := time.Since(start)

	fmt.Printf("Fibonacci(40) = %d, computed in %s\n", result, elapsed)

	// Prevent the program from exiting immediately
	select {}
}
