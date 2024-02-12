package main

import (
    "context"
)

func main() {
    ctx := context.Background()
    ctx, cancelCtx := context.WithCancel(ctx)

    go func(){

    }()
}
