package main

import "fmt"

func factorial(x int) int {
	fmt.Println(x)
	if x == 0 {
		//fmt.Println("The condition met")
		return 1
		
	}
	//t1 := x * factorial(x-1)
	//fmt.Println("calling recursive", t1)
	//fmt.Println(x)
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(4))
}
