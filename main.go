package main

import (
	"fmt"
)

func main() {
	err := printHelloWorld("Hello World.")
	if err != nil {
		fmt.Println("[main] error in printing.")
	}
}

func printHelloWorld(input string) (err error) {
	fmt.Println("Hello World")
	return err
}
