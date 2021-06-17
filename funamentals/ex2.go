package main

import "fmt"

func ex2() {
	a := (42 == 42)
	b := (41 <= 42)
	c := (43 >= 42)
	d := (44 != 42)
	e := (45 > 42)
	f := (40 < 42)
	fmt.Println("Exercise 2")
	fmt.Printf("a:%v\nb:%v\nc:%v\nd:%v\ne:%v\nf:%v\n", a, b, c, d, e, f)
}
