package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}

	{
		a := 50
		fmt.Println(a)
		b := "Try something out of the box"
		fmt.Println(b)
	}

	res := addNumber(10,20) 
	fmt.Println(res)
    //fmt.Println(b)
	// fmt.Println(y) // outside scope of y
}

func addNumber(x,y int) int {
	{
		test := 42
		fmt.Println(test)
		test2 := "This is a test"
		fmt.Println(test2)
	}

	//fmt.Println(test2)

	return x + y
}