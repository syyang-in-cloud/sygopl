package main

import "fmt"

// func reverse(s []int) []int {
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	// return s
}

func main() {
	integers := [...]int{1, 2, 3, 4, 5, 6, 7}
	s := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(reverse(integers[:]))
	reverse(integers[:])
	fmt.Println(integers[:])

	// rotate the array left by two positions
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}
