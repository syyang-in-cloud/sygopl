package main

import "fmt"

func main() {
	s := "프로그램"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println(s)
	fmt.Println(string(1234567))
}
