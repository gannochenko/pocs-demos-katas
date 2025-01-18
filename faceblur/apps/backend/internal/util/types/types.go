package types

import (
	"net/http"
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
