package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <url>");
		return;
	}

	url := os.Args[1];
	fmt.Printf("Analyzing: %s\n", url);
}