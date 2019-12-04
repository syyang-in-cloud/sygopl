package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func add(x, y int) int       { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func main() {
	x := 3
	y := 4
	fmt.Println(hypot(float64(x), float64(y)))
	fmt.Println("3+4: ", add(x, y))
	fmt.Println("3-4: ", sub(x, y))
	fmt.Println("3: ", first(x, y))
	fmt.Println("0: ", zero(x, y))

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}
