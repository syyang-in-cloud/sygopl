package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\t%x\n%x\n%x\n%t\n%T\n", "x", "X", c1, c2, c1 == c2, c1)
}
