package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "name", "Indra Mahesa")
	
	PrintName(ctx)
}

func PrintName(ctx context.Context) {
	// context.WithValue(ctx, "name", "Indra Mahesa")
	fmt.Println(ctx.Value("name"))
}