package main

import (
	"context"
	"fmt"
)

func main() {
	backgroundContext := context.Background()
	fmt.Println(backgroundContext)

	todoContext := context.TODO()
	fmt.Println(todoContext)
}