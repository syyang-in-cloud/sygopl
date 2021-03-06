package main

import (
	"fmt"
	"os"
)

func main() {
	inputString := os.Args[1:]
	for _, input := range inputString {
		fmt.Println("input: ", input, "basename: ", basename(input))
	}
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
