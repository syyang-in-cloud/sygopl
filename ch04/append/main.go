package main

import "fmt"

func appendInt(x []int, y int) []int {
	return append(x, y)
}

func main() {
	x := []int{1, 2, 3}
	y := int(4)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, y)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, y)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, y)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, y)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, y)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
	x = appendInt(x, y)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
	x = append(x, x...)
	fmt.Println(x)
	fmt.Println(x, len(x), cap(x))
}
