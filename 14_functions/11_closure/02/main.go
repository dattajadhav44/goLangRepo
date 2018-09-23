package main

import "fmt"

var x int

func main() {
	fmt.Println(increament())
	fmt.Println(increament())
}

func increament() int {
	x++
	return x
}



/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
