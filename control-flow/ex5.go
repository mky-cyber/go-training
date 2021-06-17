package main

import (
	"fmt"
)

func ex5() {
	fmt.Println("Exercise 5")
	for i := 10; i <= 100; i++ {
		fmt.Printf("For %v remainder is %v\n", i, i%4)
	}
}
