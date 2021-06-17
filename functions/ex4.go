package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "and I am", p.age, "years old.")
}

func ex4() {
	fmt.Println("Exercise 4")
	p1 := person{
		first: "John",
		last:  "Smith",
		age:   32,
	}
	p1.speak()
}
