package main 

import "fmt"

func main() {
	foo(1,2)
	foo(1,2,3)
	_aslice := []int {1,2,3,4}
	foo(_aslice...)
}

func foo(number ...int) {
	fmt.Println(number)
}


