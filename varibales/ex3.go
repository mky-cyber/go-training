package main

import "fmt"

func ex3() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("%T %T %T", x, y, z)
	fmt.Println("Exercise 3")
	fmt.Println("s:", s)
}
