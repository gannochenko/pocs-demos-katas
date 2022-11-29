package internal

import "io"

type DocumentReadHandle struct {
	Reader io.ReadCloser
}
