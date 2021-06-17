package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func ex3() {
	fmt.Println("Exercise 3")
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(s.doors)
	fmt.Println(t.doors)
}
