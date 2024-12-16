package main

import (
	"fmt"

	imagepb "api/proto/image/v1"
)

func main() {
	r := imagepb.SubmitRequest{}
	fmt.Printf("%v", r)
}
