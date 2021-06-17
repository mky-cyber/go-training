package main

import "fmt"

var x int
var g func()

func ex7() {
	fmt.Println("Exercise 7")
	f := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	f()
	fmt.Printf("%T\n", f)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	g = f
	g()
	fmt.Printf("this is g %T\n", g)

	fmt.Println("done")
}
