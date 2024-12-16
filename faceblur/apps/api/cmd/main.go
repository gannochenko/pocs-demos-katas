package main

import (
	"fmt"

	"api/internal/foo"
	imagepb "api/proto/image/v1"
)

func main() {
	r := imagepb.SubmitImageRequest{}
	fmt.Printf("%v", r)

	foo.Fn()
}
