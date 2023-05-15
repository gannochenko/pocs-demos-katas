package main

import (
	"fmt"
	"strings"

	"tree/internal/lib/tree"
)

func main() {
	materialisedTree := map[string]int32{
		"shop.foo.bar":    1,
		"shop.foo.baz":    3,
		"package.foo.bar": 7,
		"package.foo.m":   10,
	}

	treeInst := tree.New[int32]()

	for path, weight := range materialisedTree {
		treeInst.AddNode(strings.Split(path, "."), weight)
	}

	fmt.Println(treeInst.ToJSON())
}
