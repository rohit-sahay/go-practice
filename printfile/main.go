package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No filename provided")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
