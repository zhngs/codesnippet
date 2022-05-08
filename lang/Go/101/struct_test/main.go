package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	R int
}

func (p *Point) addptr() {
	p.X += 1
	p.Y += 1
}
func (p Point) add() {
	p.X += 1
	p.Y += 1
}

func main() {
	p := Point{1, 1}
	fmt.Println("p init -- ", p)
	p.add()
	fmt.Println("p add -- ", p)
	p.addptr()
	fmt.Println("p addptr -- ", p)

	q := &Point{3, 3}
	fmt.Println("q init -- ", q)
	q.add()
	fmt.Println("q add -- ", q)
	q.addptr()
	fmt.Println("q addptr -- ", q)

	var c Circle
	c.X = 1
	c.Y = 1
	c.R = 1
	fmt.Println("c init -- ", c)
	c.add()
	fmt.Println("c add -- ", c)
	c.addptr()
	fmt.Println("c add -- ", c)

	closure := Point{1, 1}
	f := closure.addptr
	f()
	fmt.Println("closure addptr -- ", closure)
}
