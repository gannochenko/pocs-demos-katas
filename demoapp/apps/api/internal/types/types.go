package types

import "net/http"

type Handler func(http.ResponseWriter, *http.Request) error
