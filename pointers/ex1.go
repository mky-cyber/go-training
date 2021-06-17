package main

import "fmt"

type person struct {
	first string
}

func ex1() {
	fmt.Println("Exercise 1")
	p1 := person{
		first: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "Miss Moneypenny"
}
