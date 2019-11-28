package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// fmt.Println("the size of bytes of data ", len(data))
		// fmt.Println("the size of bytes of byte(string(data))", len([]byte(string(data))))
		// fmt.Printf("'%x' len=%d\n", line, len(line))
		// fmt.Printf("'%x' len=%d\n", line, len(line))
		for _, line := range strings.SplitAfter(string(data), "\n") {
			if len(line) > 0 {
				counts[strings.TrimRight(line, "\n")]++
			}
		}
	}
	for line, n := range counts {
		// fmt.Printf("%d\t'%x'\n", n, line)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
