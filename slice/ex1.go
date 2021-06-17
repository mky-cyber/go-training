package main

import "fmt"

func ex1() {
	fmt.Println("Exercise 1:")
	m := [5]int{0, 1, 2, 3, 4}
	for k, v := range m {
		fmt.Printf("On index %v value is %v\n", k, v)
	}
	fmt.Printf("Type is %T\n", m)
}
