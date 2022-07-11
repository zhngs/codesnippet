package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var number int
	var floatnumber float32
	var boolean bool
	var str string
	var array [3]int
	var slice []int
	var mapzero map[string]int
	var point *int
	var function func(int) int
	var channel chan int
	var person Person
	fmt.Printf("%T zero value: %v\n", number, number)
	fmt.Printf("%T zero value: %v\n", floatnumber, floatnumber)
	fmt.Printf("%T zero value: %v\n", boolean, boolean)
	if str == "" {
		fmt.Printf("%T zero value: %v\n", str, str)
	}
	fmt.Printf("%T zero value: %v\n", array, array)
	if slice == nil {
		fmt.Printf("%T zero value: %v\n", slice, slice)
	}
	if mapzero == nil {
		fmt.Printf("%T zero value: %v\n", mapzero, mapzero)
	}
	if point == nil {
		fmt.Printf("%T zero value: %v\n", point, point)
	}
	if function == nil {
		fmt.Printf("%T zero value: %v\n", function, function)
	}
	if function == nil {
		fmt.Printf("%T zero value: %v\n", channel, channel)
	}
	fmt.Printf("%T zero value: %v\n", person, person)
	fmt.Println("")

	number2 := 6
	floatnumber2 := 3.5
	str2 := 'a'
	fmt.Printf("%T value: %v\n", number2, number2)
	fmt.Printf("%T value: %v\n", floatnumber2, floatnumber2)
	fmt.Printf("%T value: %v\n", str2, str2)

	maptest := map[string]int{}
	if maptest != nil {
		fmt.Printf("%T value: %v\n", maptest, maptest)
	}
}
