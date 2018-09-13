package main

import "fmt"

func main() {

	// you can only do this inside a func
	message := "Hello World-  Please note this sort of declaration can be done only inside func!-"
	a, b, c := 1, false, 3
	d := 4
	e := true

	fmt.Println(message, a, b, c, d, e)
}
