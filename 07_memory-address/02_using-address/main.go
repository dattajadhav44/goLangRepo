package main

import "fmt"

const meterToYards float64 = 2.09361

func main() {
	var meters float64
	fmt.Println("Enter meters swam that can happen: ")
	fmt.Scan(&meters)
	yards := meters * meterToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}

