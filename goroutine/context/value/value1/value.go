package main

import (
	"context"
	"fmt"
)

type favContextKey string

func value(ctx context.Context, k favContextKey) {
	if v := ctx.Value(k); v != nil {
		fmt.Printf("found value: %v\n", v)
		return
	}
	fmt.Printf("key not found: %v\n", k)
}

func main() {
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	value(ctx, k)
	value(ctx, favContextKey("color"))
}
