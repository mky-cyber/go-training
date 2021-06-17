package main

import "fmt"

type (
	A1 = int
	A2 = A1
)

var x1 A1

func ex4() {
	fmt.Println("Exercise 3")
	fmt.Printf("value of x: %v\n", x1)
	fmt.Printf("type of x: %T\n", x1)
	x1 = 42
	fmt.Printf("value of x: %v\n", x1)
}
