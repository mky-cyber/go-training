package main

import "fmt"

func ex1() {
	x := 42
	y := fmt.Sprintf("%b", x)
	z := fmt.Sprintf("%x", x)
	fmt.Println("Exercise 1")
	fmt.Printf("In decimal: %v\nIn binary: %v\nIn hex: %v\n", x, y, z)
}
