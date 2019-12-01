package main

import (
	"fmt"
	"math"
)

func main() {
	// page 53
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	// page 54

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("The type of x is %T, and the type of y is %T\n", x, y)
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}

	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges
	// ch03/integer/main.go:42:27:
	// invalid operation: apples + oranges (mismatched types int32 and int16)

	var compote int = int(apples) + int(oranges)
	fmt.Println(compote)

	// page 55

	f := 3.141
	ii := int(f)
	fmt.Println(f, ii)
	f = 1.99
	fmt.Println(int(f))

	f = 1e100
	ii = int(f)
	fmt.Println(f, ii)

	// page 56

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)

	// page 57
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e× = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -1, 1/z, -1/z, z/z)

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

}
