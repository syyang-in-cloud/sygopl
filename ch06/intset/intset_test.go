package intset

import (
	"fmt"
	"testing"
)

func TestExample_one(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	fmt.Println(x.String())
	fmt.Println(x)

	y.Add(9)
	fmt.Println(y.String())
}
