package main

import (
	"fmt"

	"sygopl/ch06/geometry"
)

func main() {
	perim := geometry.Path{{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	//fmt.Println(geometry.PathDistance(perim))
	fmt.Println(perim.Distance())
}
