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
		// The original problem can be understood as : https://play.golang.org/p/ov4D8jO2WHO
		// We can find more visible way with : https://play.golang.org/p/7Y7OpgGOj7w
		// Now we can see why, there is an extra split between the last "\n" and the end of string
		// the solution below can be seen easily with : https://play.golang.org/p/8cMAlENndjN

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
