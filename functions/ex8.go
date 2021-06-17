package main

import "fmt"

func ex8() {
	fmt.Println("Exercise 8")
	f := foo8()
	fmt.Println(f())
}

func foo8() func() int {
	return func() int {
		return 42
	}
}
