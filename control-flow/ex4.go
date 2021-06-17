package main

import (
	"fmt"
)

func ex4() {
	i := y
	fmt.Println("Exercise 4")
	for {
		if i > 2021 {
			break
		}
		fmt.Println(i)
		i++
	}
}
