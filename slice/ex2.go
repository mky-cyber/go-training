package main

import "fmt"

func ex2() {
	fmt.Println("Exercise 2")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for k, v := range x {
		fmt.Println(k, v)
	}
	fmt.Printf("%T\n", x)
}
