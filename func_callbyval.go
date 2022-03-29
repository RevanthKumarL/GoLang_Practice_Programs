package main

import (
	
	"fmt"
)

func swap(a, b int)int{
	var o int
	o = a
	a = b
	b = o
	return o 
}

func swap1(c, d *int)int{
	var v int
	v = *c
	*c = *d
	*d = v
	return v
}

func main() {
	var p int = 10
	var q int = 20
	fmt.Printf("p = %d and q = %d", p, q)

	swap(p, q)
	fmt.Printf("\np = %d and q = %d\n", p, q)

	var r int = 20
	var s int = 40
	fmt.Printf("\nr = %d and s = %d", r, s)

	swap1(&r, &s)
	fmt.Printf("\nr = %d and s = %d\n", r, s)
}