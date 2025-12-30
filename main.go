package main

import (
	"fmt"
	"os"

	"longurlgenerator/longurlgenerator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: long-url-generator <url>")
		fmt.Println("Example: long-url-generator https://example.com")
		os.Exit(1)
	}

	inputURL := os.Args[1]
	longURL, err := longurlgenerator.GenerateLongURL(inputURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(longURL)
}
