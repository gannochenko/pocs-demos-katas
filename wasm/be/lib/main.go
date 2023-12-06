package main

import (
	"fmt"
	"syscall/js"

	"be/internal/util"
)

func GetPerson() *util.SamplePerson {
	person := util.GetPerson()

	return person
}

// https://stackoverflow.com/questions/68798120/imports-syscall-js-build-constraints-exclude-all-go-files-in-usr-local-go-src
// https://reintech.io/blog/a-guide-to-gos-syscall-js-package-go-and-webassembly
// https://github.com/norunners/vert

type Person struct {
	Name string `js:"name"`
}

type Result struct {
	Age int `js:"age"`
}

func main() {
	js.Global().Set("greet", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) == 0 {
			return "Hello, World!"
		}
		return fmt.Sprintf("Hello, %s!", args[0].String())
	}))

	js.Global().Set("foo", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		personJSObj := args[0]

		person := Person{
			Name: personJSObj.Get("name").String(),
		}

		fmt.Printf("%v\n", person.Name)

		result := Result{
			Age: 70,
		}

		return js.ValueOf(map[string]interface{}{
			"age": result.Age,
		})
	}))

	fmt.Printf("Go executed\n")

	// block forever
	select {}
}
