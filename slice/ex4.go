package main

import "fmt"

func ex4() {
	fmt.Println("Exercise 4")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println("Append 52", x)
	x = append(x, 53, 54, 55)
	fmt.Println("Append 53 54 55", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("Append y", x)
}
