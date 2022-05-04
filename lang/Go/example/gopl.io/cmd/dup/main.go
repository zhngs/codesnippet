package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dup1()
	dup2()
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup1() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Printf("read file %s err, reason = %v\n", arg, err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func dup2() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)
		if err != nil {
			fmt.Printf("read file %s err, reason = %v\n", arg, err)
			continue
		}
		for _, lines := range strings.Split(string(data), "\n") {
			counts[lines] += 1
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
