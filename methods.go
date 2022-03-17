package main

import "fmt"

type rect struct {
	width, heigth int
}

func (r *rect) area() int {
	return r.width * r.heigth
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.heigth
}

func main() {
	r := rect{width: 10, heigth: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
