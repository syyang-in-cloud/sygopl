package main

import "fmt"

func main() {

	q := [...]int{1, 2, 3}
	r := [...]int{99: -1}

	for i, v := range q {
		fmt.Println(i, v)
	}
	for i, v := range r {
		fmt.Println(i, v)
	}
}
