package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	fmt.Printf("%T - The passed value is a integer \n", a)
	fmt.Printf("%T - The passed value is a string \n", b)
	fmt.Printf("%T - The passed value is a float \n", c)
	fmt.Printf("%T - The passed value is a bool \n", d)
	fmt.Printf("%T - The passed value is a string \n", e)
	fmt.Printf("%T - The passed value is a string \n", f)
	fmt.Printf("%T - The passed value is a int32 \n", g)
}
