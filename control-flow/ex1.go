package main

import (
	"fmt"
)

func ex1() {
	fmt.Println("Exercise 1")
	for i := 1; i <= 10000; i++ {
		fmt.Printf("i: %v\n", i)
	}
}
