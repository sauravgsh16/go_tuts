package main

import (
	"fmt"
)

type shape interface{
	getArea() float64
}

type square struct{
	sideLenght float64
}

type triangle struct{
	height float64
	base float64
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printShape(s shape) {
	fmt.Println("Area is: ", s.getArea())
}

func main() {
	s := square{
		sideLenght: 1.5,
	}
	t := triangle{
		height: 1.5,
		base: 2.5,
	}

	printShape(s)
	printShape(t)
}