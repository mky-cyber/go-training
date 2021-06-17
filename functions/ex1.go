package main

import "fmt"

func ex1() {
	fmt.Println("Exercise 1")
	n := foo()
	x, s := bar()
	fmt.Println(n, x, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big Brother is Watching"
}
