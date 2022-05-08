package main

import "fmt"

type Add interface {
	add()
}

type Addptr interface {
	addptr()
}

type Addall interface {
	Add
	Addptr
}

type Point struct {
	X, Y int
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
	var a Addall
	p := Point{1, 1}
	a = &p
	fmt.Println("a init -- ", a)
	a.add()
	fmt.Println("a add -- ", a)
	a.addptr()
	fmt.Println("a addptr -- ", a)

	// var b Addall
	// b = p //错误！p并没有实现Addall接口

	var b Add
	b = &p
	fmt.Println("b init from ptr -- ", b)
	b = p
	fmt.Println("b init -- ", b)

	var c Addptr
	c = &p
	fmt.Println("c init from ptr -- ", c)
	// c = p //错误！p没有实现Addptr接口
}
