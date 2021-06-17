package main

import "fmt"

func ex3() {
	fmt.Println("Exercise 3")
	r := foo3()
	fmt.Println("r is ", r)
}

func foo3() int {
	defer func() {
		fmt.Println("foo defer ran")
	}()
	fmt.Println("foo ran")
	fmt.Println("foo ran2")
	fmt.Println("foo ran3")
	fmt.Println("foo ran4")
	return 42
}
