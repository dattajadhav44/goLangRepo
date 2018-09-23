package main

import "fmt"

func main() {
	greater := max(2,7,9,790,309,609,788,657,555)
	fmt.Println(greater)
}

func max(number ...int) int {
	var largest int

	fmt.Println(number)
	for _, v := range number {
		if v > largest {
			largest = v 
			fmt.Println("The largest number is -", largest)
		}
	}

	return largest
}


/*
FYI
For your code to also work with only negative numbers such as

greatest := max(-200 -700)

include this as your range statement
for i, v := range numbers {
	if v > largest || i == 0 {
		largest = v
	}
}

What does that code do?

The first time through the range loop
the index, i, will be zero
so largest will be set to the first number

Originally largest is set to the zero value for an int, which is zero

Zero would be greater than any negative number

if you only have negative numbers
you need largest to be something less than zero


*/
