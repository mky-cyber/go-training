package main

import "fmt"

func ex10() {
	fmt.Println("Exercise 10")
	f := foo10()
	fmt.Println(f())
	fmt.Println(f())
	g := foo10()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
}

func foo10() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
