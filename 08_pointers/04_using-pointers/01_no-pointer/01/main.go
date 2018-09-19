package main

import "fmt"

func zero(z int) {
	fmt.Println(z) // This will print 5, as it has recieved input param
	z = 0
	fmt.Println(z) // This will print 0, as it has been reset above. 
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5, Because it is passed by value and x inside the fun
}
