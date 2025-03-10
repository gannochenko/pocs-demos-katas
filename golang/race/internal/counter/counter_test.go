package counter_test

import (
	"sync"
	"testing"

	"race/internal/counter"
)

func TestRaceCondition(t *testing.T) {
	var wg sync.WaitGroup

	// Start multiple goroutines that modify `count`
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	// No assertion is needed; we just want to detect the race condition
	t.Logf("Final count: %d", counter.GetCount())
}
