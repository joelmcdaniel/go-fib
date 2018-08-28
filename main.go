package main

import (
	"fmt"
	"go-fib/fibonacci"
)

func main() {
	fib := fibonacci.FibMemoized(5)
	fmt.Printf("%v", fib)
}
