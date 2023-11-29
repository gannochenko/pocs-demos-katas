package main

import (
	"fmt"

	"be/internal/util"
	"github.com/norunners/vert"
)

func GetPerson() *util.SamplePerson {
	person := util.GetPerson()

	return person
}

func main() {
	// https://reintech.io/blog/a-guide-to-gos-syscall-js-package-go-and-webassembly
	v := vert.ValueOf("Hello World!")

	s := ""
	v.AssignTo(&s)

	constName := util.GetSampleConstant()
	fmt.Printf("Hello, WebAssembly: %s!\n", constName)
}
