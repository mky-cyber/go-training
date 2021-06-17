package main

import (
	"fmt"
)

var y1 int

func ex5() {
	fmt.Println("Exercise 5")
	fmt.Println("variable x1:", x1)
	fmt.Printf("x1 type: %T\n", x1)
	x1 = 53
	fmt.Println("variable x1:", x1)
	y1 = int(x1)
	fmt.Println("variable y1:", y1)
	fmt.Printf("y1 type: %T\n", y1)
}
