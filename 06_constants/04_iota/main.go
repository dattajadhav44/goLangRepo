package main

import "fmt"

const (
	a = iota // 0
	b        // 1
	c        // 2
	d        // 3 these automaticaly increase the rest of the sequences. 
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
