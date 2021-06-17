package main

import (
	"fmt"
)

func ex7() {
	fmt.Println("Exercise 7")
	for i := 0; i < 20; i++ {
		if i%3 == 1 {
			fmt.Printf("For %v remainder is one\n", i)
		} else if i%3 == 2 {
			fmt.Printf("For %v remainder is two\n", i)
		} else {
			fmt.Printf("For %v remainder is zero\n", i)
		}
	}
}
