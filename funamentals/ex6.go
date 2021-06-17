package main

import (
	"fmt"
)

const (
	c = 2021 + iota
	d = 2021 + iota
	e = 2021 + iota
	f = 2021 + iota
)

func ex6() {
	fmt.Println("Exercise 6")
	fmt.Println("Next 4 years", c, d, e, f)
}
