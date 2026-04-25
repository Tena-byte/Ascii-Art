package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . input.txt")
		return
	}
	file := os.Args[2]

	file1, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("error")
	}
	output := string(file1)
	fmt.Println(output)
}
