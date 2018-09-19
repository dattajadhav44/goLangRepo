package main

import "fmt"

func main() {

	a := 55

	fmt.Println(a)  // This will print 55
	fmt.Println(&a) // This will print a's Address

	b := &a
	fmt.Println(b)  // this should print the a's address
	fmt.Println(*b) // this should display the the value that is 55 at a memory location
	*b = 65         // this should change the value to a and b to 65
	fmt.Println(a)  // should display the 65
	fmt.Println(*b)

    // this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
