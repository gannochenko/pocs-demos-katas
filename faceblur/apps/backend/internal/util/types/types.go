package types

import (
	"math"
	"net/http"
	"sync"
	"time"
)

type CbWithError func() error
type Cb func()

type HTTPHandler func(http.ResponseWriter, *http.Request) error

func Unwrap[T any](input []T, def T) T {
	if input == nil || len(input) == 0 {
		return def
	}

	return input[0]
}

func GetNowUTC() time.Time {
	return time.Now().UTC().Truncate(time.Second)
}

func CloseChannelSafely[P any](ch chan P) {
	select {
	case _, ok := <-ch:
		if !ok {
			// Channel is already closed
			return
		}
	default:
		// Channel is not yet closed; safe to close
	}
	close(ch)
}

func GetSyncMapSize(m *sync.Map) int {
	count := 0
	m.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	return count
}

func Float32ToInt32(value float32) int32 {
	return int32(math.Round(float64(value)))
}