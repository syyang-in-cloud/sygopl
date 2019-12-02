package main

import (
	"fmt"
	"math/rand"
)

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func Sort(values []int) {
	var root *tree
	// form the tree based on the given values
	for _, v := range values {
		root = add(root, v)
	}
	// pass the empty value slice
	// so that the appendValues can return sorted values
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	values := [100]int{}
	for i := range values {
		values[i] = rand.Intn(100)
	}
	fmt.Println(values)
	Sort(values[:])
	fmt.Println(values)

}
