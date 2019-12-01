package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputString := os.Args[1:]
	for _, input := range inputString {
		fmt.Println("input: ", input, "basename: ", basename(input))
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
