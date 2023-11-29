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

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		fourWheel: false,
	}

	car1 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "silver",
		},
		luxury: true,
	}

	fmt.Printf("%+v\n", t1)
	fmt.Printf("%+v\n", car1)

	fmt.Printf("%v\n", t1.color)
	fmt.Printf("%v\n", car1.color)
}
