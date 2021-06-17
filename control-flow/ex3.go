package main

import (
	"fmt"
)

var y int = 1993

func ex3() {
	fmt.Println("Exercise 3")
	i := y
	for i <= 2021 {
		fmt.Println(i)
		i++
	}
}
