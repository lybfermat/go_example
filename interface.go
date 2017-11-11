package main 

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect)area() float64 {
	return r.width * r.height
}

func (r *rect)perim()float64 {
	return 2 * r.width + 2 * r.height
}

func (r *circle)area()float64 {
	return 3.14 * r.radius * r.radius
}

func (r *circle)perim() float64 {
	return 3.14 * r.radius * 2
}

func show(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width:10.2, height:30}
	c := circle{radius:5}
	show(&r)
	show(&c)
}