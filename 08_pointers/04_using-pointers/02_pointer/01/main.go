package main

import "fmt"

func zero(z *int) {
	fmt.Println(z)
	fmt.Println(&z)
	*z = 0       // the value of x has been changed becuase of the memory pointer
	fmt.Println(z)
	fmt.Println(&z)

}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
