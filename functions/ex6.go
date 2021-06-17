package main

import "fmt"

func ex6() {
	fmt.Println("Exercise 6")
	func() {
		for i := 0; i < 50; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}
