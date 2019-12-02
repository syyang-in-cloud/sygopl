package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func main() {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"x", "y", "z"}

	fmt.Println(k(list1))
	fmt.Println(k(list2))

	Add(list1)
	Add(list2)
	fmt.Println(Count(list1))
	fmt.Println(Count(list2))

	for k, v := range m {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
