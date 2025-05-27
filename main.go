package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}
