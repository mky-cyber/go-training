package main

import "fmt"

func ex2() {
	fmt.Println("Exercise 2")
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo2(ii...)
	fmt.Println(n)
	n2 := bar2(ii)
	fmt.Println(n2)
}

func foo2(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar2(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
