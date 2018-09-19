package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
	d = iota // 3 

	// iota increases the number, infact it keeps increasing the number
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
