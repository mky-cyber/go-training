package main

import (
	"fmt"
)

func ex4() {
	fmt.Println("Exercise 4")
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)
}
