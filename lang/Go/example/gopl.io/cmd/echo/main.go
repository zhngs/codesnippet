package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, seq := "", ""
	for _, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echo1()
	echo2()
	echo3()
}
