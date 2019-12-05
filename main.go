package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	x := add(5, 3)
	fmt.Println("x: ", x)
}

func add(a, b int) int {
	return a + b
}
