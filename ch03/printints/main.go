package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func myIntsToString(values []int) string {
	var buf string
	buf = fmt.Sprintf("[")
	for i, v := range values {
		if i > 0 {
			buf += ", "
		}
		buf += fmt.Sprintf("%d", v)
	}
	buf += "]"
	return buf
}

func main() {
	fmt.Println(myIntsToString([]int{1, 2, 3}))
}
