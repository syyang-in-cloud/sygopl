package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	other := map[string]int{
		"alice":   31,
		"charlie": 32,
	}
	fmt.Println("The ages looks ", ages)
	fmt.Println("The other looks ", other)

	fmt.Println("alice was deleted from ages map")
	delete(ages, "alice")

	fmt.Println("The ages looks ", ages)
	fmt.Println("The other looks ", other)

	for name, age := range other {
		fmt.Println(name, age)
	}

	var names []string
	for name := range other {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, other[name])
	}

	fmt.Println("ages and other is same: ", equal(ages, other))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
