package main

import "fmt"

func main() {


	increamentA := wrapper()
	increamentB := wrapper()

	fmt.Println("A :", increamentA())
	fmt.Println("A :", increamentA())
	fmt.Println("B :", increamentB())
	fmt.Println("B :", increamentB())
	fmt.Println("B :", increamentB())
}

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
